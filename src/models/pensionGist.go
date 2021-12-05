package models

import (
	// "database/sql"
	"gopkg.in/guregu/null.v4"
)

//the null.Int is to handle null values from the database

type PensionGist struct {
	ID                  null.Int
	TITLE               null.String
	HEADER_IMAGE        null.String
	BODY_IMAGE          null.String
	VIEWS               null.Int
	LIKES               null.Int
	DATE_POSTED         null.String
	CONTENT             null.String
	RELATED_ARTICLE_ID1 null.Int
	RELATED_ARTICLE_ID2 null.Int
	RELATED_ARTICLE_ID3 null.Int
	RELATED_ARTICLE_ID4 null.Int
	STATUS              null.String
}
