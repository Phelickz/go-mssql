package routes

import "github.com/gin-gonic/gin"

func InitRoutes(r *gin.Engine) {
	AccessCredentials(r)
	PensionGistRoute(r)
	TokenRoute(r)
}
