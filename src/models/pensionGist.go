package models

import "database/sql"

type PensionGist struct {
	ID                  int
	TITLE               string
	HEADER_IMAGE        string
	BODY_IMAGE          string
	VIEWS               int
	LIKES               int
	DATE_POSTED         string
	CONTENT             string
	RELATED_ARTICLE_ID1 int
	RELATED_ARTICLE_ID2 int
	RELATED_ARTICLE_ID3 int
	RELATED_ARTICLE_ID4 int
	STATUS              sql.NullString
}
