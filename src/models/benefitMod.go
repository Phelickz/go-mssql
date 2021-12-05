package models

import "gopkg.in/guregu/null.v4"

type Benefit struct {
	PIN               null.String `json:"pin"`
	BENEFIT_TYPE      null.String `json:"benefit_type"`
	MONTHLY_HOUSING   null.String `json:"monthly_housing"`
	MONTHLY_TRANSPORT null.String `json:"monthly_transport"`
	MONTHLY_BASIC     null.String `json:"monthly_basic"`
	CURRENT_EMAIL     null.String `json:"current_email"`
	CURRENT_MOBILE    null.String `json:"current_mobile"`
	DOCUMENT_CONSENT  null.String `json:"document_consent"`
}
