package models

type Localize struct {
	ID    int `gorm:"primaryKey,autoIncrement"`
	Key   string
	Value string
}

func (Localize) TableName() string {
	return "localize"
}
