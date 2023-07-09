package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var DB *gorm.DB

func ConnectDB () { 
	var err error

	sqlServer := os.Getenv("SQL_SERVER")
	sqlUser := os.Getenv("SQL_USER")
	sqlPassword := os.Getenv("SQL_PASSWORD")
	sqlDatabase := os.Getenv("SQL_DATABASE")
	sqlPort := os.Getenv("SQL_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", sqlServer, sqlUser, sqlPassword, sqlDatabase, sqlPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Connection to Database")
	}

	
}