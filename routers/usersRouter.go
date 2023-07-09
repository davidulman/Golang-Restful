package routers

import (
	"github.com/davidulman/Golang-Restful/controllers"
	"github.com/gin-gonic/gin"
)




func UserRoutes (r* gin.RouterGroup) {

	r.POST("", controllers.CreateUser)


}