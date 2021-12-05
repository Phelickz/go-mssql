package routes

import (
	"github.com/Phelickz/go-sql/src/controllers"
	"github.com/gin-gonic/gin"
)

func TokenRoute(r *gin.Engine) {
	r.GET("/get-token", controllers.SaveTokenInDB())
}
