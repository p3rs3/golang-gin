package model

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Order struct {
	Id        int       `gorm:"type:int;primary_key" json:"id" binding:"required"`
	Price     int       `gorm:"type:int" json:"price" binding:"required"`
	Name      string    `gorm:"type:character varying(255)" json:"name" binding:"required"`
	Discount  null.Int  `gorm:"type:int;default:null" json:"discount" binding:"required"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:NOW()" json:"created_at" binding:"required"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:NOW()" json:"updated_at" binding:"required"`
	UserId    int       `json:"-"`
}
