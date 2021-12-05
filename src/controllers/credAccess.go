package controllers

import (
	"context"
	"database/sql"

	// "database/sql"

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
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()

		result, err := dbInstance.QueryContext(ctx, "getPensionsGist")
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			// c.Abort()
		}
		gists := []models.PensionGist{}

		for result.Next() {
			var ID int
			var TITLE string
			var HEADER_IMAGE string
			var BODY_IMAGE string
			var VIEWS int
			var LIKES int
			var DATE_POSTED string
			var CONTENT string
			var RELATED_ARTICLE_ID1 int
			var RELATED_ARTICLE_ID2 int
			var RELATED_ARTICLE_ID3 int
			var RELATED_ARTICLE_ID4 int
			var STATUS sql.NullString
			err2 := result.Scan(&ID, &TITLE, &HEADER_IMAGE, &BODY_IMAGE, &VIEWS, &LIKES, &DATE_POSTED, &CONTENT, &RELATED_ARTICLE_ID1, &RELATED_ARTICLE_ID2, &RELATED_ARTICLE_ID3, &RELATED_ARTICLE_ID4, &STATUS)
			if err2 != nil {
				log.Panic(err2)
			} else {
				gist := models.PensionGist{ID: ID, TITLE: TITLE, HEADER_IMAGE: HEADER_IMAGE, BODY_IMAGE: BODY_IMAGE, VIEWS: VIEWS, LIKES: LIKES, DATE_POSTED: DATE_POSTED, CONTENT: CONTENT, RELATED_ARTICLE_ID1: RELATED_ARTICLE_ID1, RELATED_ARTICLE_ID2: RELATED_ARTICLE_ID2, RELATED_ARTICLE_ID3: RELATED_ARTICLE_ID3, RELATED_ARTICLE_ID4: RELATED_ARTICLE_ID4}
				gists = append(gists, gist)
				c.JSON(200, gists)
			}

			// fmt.Println(err)
		}

		dbInstance.Close()
		// c.JSON(http.StatusOK, result)

	}
}

// var account = "abc"
// _, err := db.ExecContext(ctx, "sp_RunMe",
// 	sql.Named("ID", 123),
// 	sql.Named("Account", sql.Out{Dest: &account}),
// )

// sql.Named("pin", "PEN100599222817"), sql.Named("password", "12345")
