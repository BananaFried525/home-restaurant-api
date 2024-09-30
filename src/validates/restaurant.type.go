package validates

type GetTableRequest struct {
}

type AddTableRequest struct {
	TableNumber    int `json:"table_number" binding:"required" min:"1"`
	CustomerNumber int `json:"customer_number" binding:"required" min:"1"`
}

type GetMenuRequest struct {
	TableNumber int    `form:"table_number" binding:"required" min:"1"`
	TableID     string `form:"table_id" binding:"required"`
}

type FoodOrderAttribute struct {
	FoodID  int    `json:"food_id" binding:"required"`
	Amount  int    `json:"amount" binding:"required" min:"1"`
	Comment string `json:"comment"`
}
type AddOrderRequest struct {
	TableNumber int                  `json:"table_number" binding:"required" min:"1"`
	TableID     string               `json:"table_id" binding:"required"`
	Food        []FoodOrderAttribute `json:"food" binding:"required"`
}
