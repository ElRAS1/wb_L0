package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

type Store struct {
	config *Config
	db     *sqlx.DB
	Cache  map[string]interface{}
}

func New(config *Config) *Store {
	return &Store{
		config: config,
		Cache:  make(map[string]interface{}),
	}
}

func (s *Store) Open() error {
	DatabaseUrl := fmt.Sprintf("host=%v dbname=%v user=%v password=%v sslmode=disable", s.config.DBHost, s.config.DBName, s.config.DBUser, s.config.DBPassword)
	db, err := sqlx.Connect("postgres", DatabaseUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Store) NatsSubscribe(ns stan.Conn) {

	ns.Subscribe("wb", func(msg *stan.Msg) {
		data := Order{}
		err := json.Unmarshal(msg.Data, &data)

		if err != nil {
			log.Fatalln(err)
		}
		err = s.insertOrder(s.db, data)
		if err != nil {
			log.Println(err)
			return
		}

	}, stan.DurableName("wb"))

}

func (s *Store) NatsPublish(ns stan.Conn) error {
	data, err := s.jsonData()

	if err != nil {
		return err
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error serializing Order to JSON: %v", err)
	}

	err = ns.Publish("wb", []byte(jsonData))
	if err != nil {
		log.Fatalf("Error while trying to send msg: %v", err)
	}

	return nil
}

func (s *Store) jsonData() (*Order, error) {
	var order Order
	file, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(file, &order)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	if err != nil {
		log.Fatalln(err)
	}

	return &order, nil
}

func (s *Store) insertOrder(db *sqlx.DB, order Order) error {
	// Вставка данных в таблицу orders
	query := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING order_uid`
	_, err := db.Exec(query, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	// Вставка данных в таблицу deliveries
	deliveryQuery := `INSERT INTO deliveries (order_id, client_name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.Exec(deliveryQuery, order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}

	// Вставка данных в таблицу payments
	paymentQuery := `INSERT INTO payments (order_id, transaction_id, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = db.Exec(paymentQuery, order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}

	// Вставка данных в таблицу items
	for _, item := range order.Items {
		itemQuery := `INSERT INTO items (order_id, chrt_id, track_number, price, rid, item_name, sale, item_size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
		_, err = db.Exec(itemQuery, order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			return err
		}
	}

	s.Cache[order.OrderUID] = order

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
