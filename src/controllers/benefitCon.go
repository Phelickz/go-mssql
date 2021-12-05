package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/Phelickz/go-sql/src/models"
	"github.com/gin-gonic/gin"
)

func InitiateBenefit() gin.HandlerFunc {
	return func(c *gin.Context) {

		var benefit models.Benefit

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		//parsing data from request body
		err := c.BindJSON(&benefit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
			c.Abort()
		}

		_, dbErr := dbInstance.ExecContext(ctx,
			"addInitiatedRequest",
			sql.Named("pin", benefit.PIN),
			sql.Named("benefit_type", benefit.BENEFIT_TYPE),
			sql.Named("monthly_housing", benefit.MONTHLY_HOUSING),
			sql.Named("monthly_transport", benefit.MONTHLY_TRANSPORT),
			sql.Named("monthly_basic", benefit.MONTHLY_BASIC),
			sql.Named("current_email", benefit.CURRENT_EMAIL),
			sql.Named("current_mobile", benefit.CURRENT_MOBILE),
			sql.Named("document_consent", benefit.DOCUMENT_CONSENT))

		if dbErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": dbErr.Error()})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{"message": "Update successful"})

	}
}
