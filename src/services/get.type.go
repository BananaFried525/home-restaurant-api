package services

type GetTableInfoParams struct {
	Limit  int `form:"limit,default=10"`
	Offset int `form:"offset,default=0"`
}
