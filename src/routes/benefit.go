package routes

import (
	"github.com/Phelickz/go-sql/src/controllers"
	"github.com/gin-gonic/gin"
)

func BenefitRoutes(r *gin.Engine) {
	r.POST("/initiate-request", controllers.InitiateBenefit())
}
