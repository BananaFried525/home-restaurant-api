package boots

import "gorm.io/gorm"

type Boot struct {
	Database *gorm.DB
	Config   *ConfigsAttribute
}

func New() *Boot {
	config := InitConfig()
	db := InitDatabase()

	return &Boot{
		Database: db,
		Config:   config,
	}
}
