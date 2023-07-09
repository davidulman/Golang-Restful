package main

import (
	"github.com/davidulman/Golang-Restful/initializers"
	"github.com/davidulman/Golang-Restful/routers"
	"github.com/gin-gonic/gin"
)


func init () { 
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.AutoMigrate()

}

func main () {


	r := gin.Default()



	api :=r.Group("/api/v1")

	
	routers.UserRoutes(api.Group("/users"))




	r.Run() 

}