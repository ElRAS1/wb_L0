package store

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
)

func (s *Store) checkData(data *Order) bool {
	if _, ok := s.Cache[data.OrderUID]; ok {
		return true
	}

	var order Order
	query := `SELECT * FROM orders WHERE order_uid = $1;`

	err := s.db.QueryRowx(query, data.OrderUID).StructScan(&order)

	if err != nil {
		if err == sql.ErrNoRows {
			err = s.insertOrder(s.db, *data)

			if err != nil {
				log.Fatalln(err)
			}
			return false
		}
		log.Fatalln(err)

	}
	s.CacheMutex.Lock()
	s.Cache[data.OrderUID] = data
	s.CacheMutex.Unlock()
	return false
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
