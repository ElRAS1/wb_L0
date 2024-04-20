package main

type Item struct {
	ID          int    `db:"id" json:"item_id,omitempty"`
	OrderID     string `db:"order_id" json:"order_id,omitempty"`
	ChrtID      int    `db:"chrt_id" json:"chrt_id"`
	TrackNumber string `db:"track_number" json:"track_number"`
	Price       int    `db:"price" json:"price"`
	Rid         string `db:"rid" json:"rid"`
	Name        string `db:"item_name" json:"name"`
	Sale        int    `db:"sale" json:"sale"`
	Size        string `db:"item_size" json:"size"`
	TotalPrice  int    `db:"total_price" json:"total_price"`
	NmID        int    `db:"nm_id" json:"nm_id"`
	Brand       string `db:"brand" json:"brand"`
	Status      int    `db:"status" json:"status"`
}
