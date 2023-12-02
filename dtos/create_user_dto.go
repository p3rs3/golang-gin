package dto

type CreateUserDto struct {
	Name       string `json:"name" binding:"required,min=1,max=30"`
	Secondname string `json:"secondname" binding:"required,min=1,max=30"`
	Age        int    `json:"age" binding:"required,min=1,max=99"`
}
