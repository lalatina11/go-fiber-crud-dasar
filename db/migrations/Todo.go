package migrations

import (
	"fiber-go/db"
	"fiber-go/db/models"
	"fmt"
)

// migrate the todo table

func TodoMigration() {
	err := db.DB().AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Todo Table Migration completed successfully")
}
