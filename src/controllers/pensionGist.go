package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Phelickz/go-sql/src/models"
	"github.com/gin-gonic/gin"
)

// var dbInstance = database.OpenDb()

func FetchPensionGist() gin.HandlerFunc {
	return func(c *gin.Context) {

		//creating a context
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		//calling stored procedure
		result, err := dbInstance.QueryContext(ctx, "getPensionsGist")

		//checking for an error
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			// c.Abort()
		}

		//creating a list of PensionGist structs
		gists := []models.PensionGist{}

		//looping over the results and parsing the data
		for result.Next() {
			var pension models.PensionGist
			err2 := result.Scan(&pension.ID, &pension.TITLE, &pension.HEADER_IMAGE, &pension.BODY_IMAGE, &pension.VIEWS, &pension.LIKES, &pension.DATE_POSTED, &pension.CONTENT, &pension.RELATED_ARTICLE_ID1, &pension.RELATED_ARTICLE_ID2, &pension.RELATED_ARTICLE_ID3, &pension.RELATED_ARTICLE_ID4, &pension.STATUS)
			if err2 != nil {
				log.Panic(err2)
				c.JSON(http.StatusInternalServerError, gin.H{"Error": err2})
			} else {
				gists = append(gists, pension)
				c.JSON(http.StatusOK, gists)
			}

			// fmt.Println(err)
		}

		dbInstance.Close()
		// c.JSON(http.StatusOK, result)

	}
}
