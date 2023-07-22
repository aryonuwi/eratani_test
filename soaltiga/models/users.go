package models

type Users struct {
	ID             int64  `json:"id";primaryKey`
	Country        string `json:"country" gorm:"type:varchar(255)"`
	CreditCardType string `json:credit_card_type gorm:"type:varchar(255)"`
	CreditCard     string `json:credit_card gorm:"type:varchar(255)"`
	FirstName      string `json:first_name gorm:"type:varchar(255)"`
	LastName       string `json:last_name gorm:"type:varchar(255)"`
}

type RegsiterdUsers struct {
	Country        string `json:"country" gorm:"type:varchar(255)"`
	CreditCardType string `json:credit_card_type gorm:"type:varchar(255)"`
	CreditCard     string `json:credit_card gorm:"type:varchar(255)"`
	FirstName      string `json:first_name gorm:"type:varchar(255)"`
	LastName       string `json:last_name gorm:"type:varchar(255)"`
}
