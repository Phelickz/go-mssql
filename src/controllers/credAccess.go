package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Phelickz/go-sql/src/database"
	"github.com/Phelickz/go-sql/src/models"
	"github.com/gin-gonic/gin"
)

var dbInstance = database.OpenDb()

func CredAccess() gin.HandlerFunc {
	return func(c *gin.Context) {

		//getting pin from query params
		pin, _ := c.GetQuery("pin")
		fmt.Println(pin)

		if pin == "" {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Pin must be provided"})
		}

		//creating context
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		//running stored procedure
		result, err := dbInstance.QueryContext(ctx, "GetAccessCredentials", sql.Named("pin", pin))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			// c.Abort()
		}
		defer result.Close()

		// creating a list of user structs
		users := []models.User{}

		//parsing
		for result.Next() {
			var user models.User
			err2 := result.Scan(&user.ID, &user.PIN, &user.SURNAME, &user.FIRSTNAME, &user.OTHERNAMES, &user.PASSWORD, &user.STATUS, &user.EMAIL, &user.MOBILE_PHONE, &user.DEVICE_ID)
			if err2 != nil {
				log.Panic(err2)
				c.JSON(http.StatusInternalServerError, err2)
			} else {
				users = append(users, user)
			}
		}

		//checking for empty response and sending response to user
		if len(users) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "User does not exist on the database"})
		} else {
			c.JSON(http.StatusOK, users)
		}
	}
}

func SetPasswordAndDeviceID() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		pin, _ := c.GetQuery("pin")
		password, _ := c.GetQuery("password")
		device_id, _ := c.GetQuery("device_id")

		_, err := dbInstance.ExecContext(ctx, "UpdateAccessCredentials", sql.Named("pin", pin), sql.Named("password", password), sql.Named("device_id", device_id))

		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{"message": "Update Successful"})
	}
}
