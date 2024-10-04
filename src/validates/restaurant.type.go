package validates

type GetTableListRequest struct {
	Limit  int `form:"limit,default=10"`
	Offset int `form:"offset,default=0"`
}

type CreateTableOrderRequest struct {
	TableInfoID uint `json:"table_info_id" binding:"required"`
}
