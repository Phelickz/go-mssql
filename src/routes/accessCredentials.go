package routes

import (
	"github.com/Phelickz/go-sql/src/controllers"
	"github.com/gin-gonic/gin"
)

func AccessCredentials(r *gin.Engine) {
	r.GET("/activate-pin", controllers.CredAccess())
	r.GET("/set-password", controllers.SetPasswordAndDeviceID())
}
