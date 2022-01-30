package migration

import (
	"fmt"
	"log"

	"github.com/alfaroqi/learning-go-fiber-with-gorm/database"
	"github.com/alfaroqi/learning-go-fiber-with-gorm/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}) // AutoMigrate will create table for you if not exist
	if err != nil {
		log.Println("Error di :", err)
	}

	fmt.Println("Migration Successful")
}
