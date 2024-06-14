package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	var err error

	DSN := "host=localhost user=mariovlv password=221211ab0 dbname=gorm port=5432"

	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connected")
}
