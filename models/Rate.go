package models

import (
	"time"

	"github.com/google/uuid"
)

type Rate struct {
	ID               uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;default:gen_random_uuid()"`
	Srno             int       `json:"srno" gorm:"int;not null"`
	SellingRate      float32   `json:"selling_rate"`
	BuyingRate       float32   `json:"buying_rate"`
	AverageRate      float32   `json:"average_rate"`
	MidRate          float32   `json:"mid_rate"`
	CurrencyID       uuid.UUID `json:"currency_id" gorm:"type:uuid;not null"`
	CreatedDate      time.Time `json:"created_date" gorm:"type:date;not null;default:current_date"`
	CreatedTime      string    `json:"created_time" gorm:"default:to_char(now(),'HH24:MI:SS AM')"`
	CreatedbyID      uuid.UUID `json:"createdby_id" gorm:"type:uuid;not null"`
	LastmodifiedDate string    `json:"lastmodified_date"`
	LastmodifiedTime string    `json:"lastmodified_time" gorm:"type:varchar(12);null"`
	LastmodifiedbyID uuid.UUID `json:"lastmodifiedby_id" gorm:"type:uuid;default:null"`
	IsApproved       bool      `json:"is_approved"`
	ApprovedbyID     string    `json:"approvedby_id" gorm:"type:uuid;null;default:null"`
	ApprovedTime     string    `json:"approved_time"`
	RejectedDate     string    `json:"rejected_date"`
	RejectedTime     string    `json:"rejected_time" gorm:"type:varchar(12);null"`
	RejectedbyID     uuid.UUID `json:"rejectedby_id" gorm:"type:uuid;null;default:null"`

	FkCurrencyID       Currency `gorm:"foreignKey:CurrencyID;"`
	FkCreatedbyID      User     `gorm:"foreignKey:CreatedbyID;"`
	FkApprovedbyID     User     `gorm:"foreignKey:ApprovedbyID;"`
	FkRejectedbyID     User     `gorm:"foreignKey:RejectedbyID;"`
	FkLastmodifiedbyID User     `gorm:"foreignKey:LastmodifiedbyID;"`
}
