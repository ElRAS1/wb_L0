package main

type Payment struct {
	TransactionID string `db:"transaction_id" json:"transaction"`
	OrderID       string `db:"order_id" json:"order_id,omitempty"`
	RequestID     string `db:"request_id" json:"request_id"`
	Currency      string `db:"currency" json:"currency"`
	Provider      string `db:"provider" json:"provider"`
	Amount        int    `db:"amount" json:"amount"`
	PaymentDT     int    `db:"payment_dt" json:"payment_dt"`
	Bank          string `db:"bank" json:"bank"`
	DeliveryCost  int    `db:"delivery_cost" json:"delivery_cost"`
	GoodsTotal    int    `db:"goods_total" json:"goods_total"`
	CustomFee     int    `db:"custom_fee" json:"custom_fee"`
}
