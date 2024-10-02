package validates

type GetTableRequest struct {
	Limit  int `form:"limit,default=10"`
	Offset int `form:"offset,default=0"`
}

type AddTableRequest struct {
	TableNumber    uint `json:"table_number" binding:"required" min:"1"`
	CustomerNumber int  `json:"customer_number" binding:"required" min:"1"`
}

type GetMenuRequest struct {
}

type GetTableDetailRequest struct {
	TableNumber uint `form:"table_number" binding:"required" min:"1"`
}

type FoodOrderAttribute struct {
	FoodID  uint    `json:"food_id" binding:"required"`
	Amount  int     `json:"amount" binding:"required" min:"1"`
	Comment *string `json:"comment"`
}
type OrderFoodRequest struct {
	TableNumber    uint                 `json:"table_number" binding:"required" min:"1"`
	ReserveTableID uint                 `json:"reserve_table_id" binding:"required" min:"1"`
	Food           []FoodOrderAttribute `json:"food" binding:"required" min:"1"`
}
