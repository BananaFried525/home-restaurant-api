package controllers

const (
	FoodActive          FoodStatus = "active"
	FoodTemporaryRunout FoodStatus = "temporary_runout"
	FoodInactive        FoodStatus = "inactive"
)

type FoodStatus string

type FoodAttribute struct {
	ID                uint
	Name              string
	Description       string
	DisplayImage      string
	Price             int
	VatPercent        int
	ServiceFeeFixed   int
	ServiceFeePercent int
	Status            FoodStatus
}

type MenuAttribute struct {
	FoodID          uint
	FoodName        string
	FoodPrice       int
	FoodImage       string
	FoodDescription string
	FoodStatus      string
}

type TableAttribute struct {
	ID             string
	TableNumber    int
	CustomerNumber int
	Orders         []OrderAttribute
	StartDate      string
	EndDate        string
}

const (
	OrderPending OrderStatus = "pending"
	OrderDone    OrderStatus = "Done"
)

type OrderStatus string
type OrderAttribute struct {
	ID          string
	TableID     string
	FoodID      int
	Status      OrderStatus
	ConfirmDate string
	FinishDate  string
}
