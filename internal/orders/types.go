package orders

type OrderProduct struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Quantity int32  `json:"quantity"`
	Price    int32  `json:"price"`
}

type createOrderParams struct {
	CustomerID int64          `json:"customer_id"`
	Products   []OrderProduct `json:"products"`
}
