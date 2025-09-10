package migrations

import (
	"fmt"

	"github.com/lalatina11/go-fiber-crud-dasar/db"
	"github.com/lalatina11/go-fiber-crud-dasar/db/models"
)

// migrate the todo table

func TodoMigration() {
	err := db.DB().AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Todo Table Migration completed successfully")
}
