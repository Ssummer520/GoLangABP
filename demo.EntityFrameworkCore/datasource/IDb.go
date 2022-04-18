package datasource

import "gorm.io/gorm"

type IDb interface {
	DB() *gorm.DB
	Connect() error
}
