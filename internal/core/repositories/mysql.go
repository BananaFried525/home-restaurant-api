package repositories

import "gorm.io/gorm"

type DB struct {
	db *gorm.DB
}

func NewMySQL(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}
