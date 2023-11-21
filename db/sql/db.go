package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {
	dsn := "root:souviks72@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to db: %+v\n", err)
		return &Database{}, err
	}

	return &Database{Client: db}, nil
}
