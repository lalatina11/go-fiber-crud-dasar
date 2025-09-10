package db

import (
	"fmt"
	"log"

	"github.com/lalatina11/go-fiber-crud-dasar/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	dsn := lib.EnvDatabaseUrl()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	return db
}
