package initializers

import (
	"github.com/davidulman/Golang-Restful/models"
)



func AutoMigrate () {

DB.AutoMigrate(&models.User{})



}