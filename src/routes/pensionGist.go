package routes

import (
	"github.com/Phelickz/go-sql/src/controllers"
	"github.com/gin-gonic/gin"
)

func PensionGistRoute(r *gin.Engine) {
	r.GET("/pension-gists", controllers.FetchPensionGist())
}
