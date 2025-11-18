package products

type createProductParams struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}
