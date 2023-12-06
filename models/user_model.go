package model

type User struct {
	Id         int    `gorm:"type:int;primary_key" json:"id" binding:"required"`
	Name       string `gorm:"type:character varying(255)" json:"name" binding:"required"`
	Secondname string `gorm:"type:character varying(255)" json:"secondname" binding:"required"`
	Age        int    `gorm:"type:int" json:"age" binding:"required"`
}
