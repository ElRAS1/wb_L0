package main

type Delivery struct {
	ID      int    `db:"id" json:"delivery_id,omitempty"`
	OrderID string `db:"order_id" json:"order_id,omitempty"`
	Name    string `db:"client_name" json:"name"`
	Phone   string `db:"phone" json:"phone"`
	Zip     string `db:"zip" json:"zip"`
	City    string `db:"city" json:"city"`
	Address string `db:"address" json:"address"`
	Region  string `db:"region" json:"region"`
	Email   string `db:"email" json:"email"`
}
