package model

type User struct {
	Id         int    `gorm:"type:int;primary_key"`
	Name       string `gorm:"type:character varying(255)"`
	Secondname string `gorm:"type:character varying(255)"`
	Age        int    `gorm:"type:int"`
}
