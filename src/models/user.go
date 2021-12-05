package models

import "gopkg.in/guregu/null.v4"

type User struct {
	ID           null.String
	PIN          null.String
	SURNAME      null.String
	FIRSTNAME    null.String
	OTHERNAMES   null.String
	PASSWORD     null.String
	STATUS       null.String
	EMAIL        null.String
	MOBILE_PHONE null.String
	DEVICE_ID    null.String
}
