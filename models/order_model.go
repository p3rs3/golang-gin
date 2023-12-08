package model

type Order struct {
	Id        int    `gorm:"type:int;primary_key" json:"id" binding:"required"`
	Price     string `gorm:"type:int" json:"price" binding:"required"`
	Name      string `gorm:"type:character varying(255)" json:"name" binding:"required"`
	Discount  int    `gorm:"type:int;default:null" json:"discount" binding:"required"`
	CreatedAt int    `gorm:"type:timestamptz;default:NOW()" json:"created_at" binding:"required"`
	UpdatedAt int    `gorm:"type:timestamptz;default:NOW()" json:"updated_at" binding:"required"`
}
