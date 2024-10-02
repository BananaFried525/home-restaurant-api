package models

type Localize struct {
	ID    uint `gorm:"primaryKey;autoIncrement:true"`
	Key   string
	Value string
}

func (Localize) TableName() string {
	return "localize"
}
