package dto

type UpdateUserDto struct {
	Name       string `json:"name" binding:"omitempty,min=1,max=30"`
	Secondname string `json:"secondname" binding:"omitempty,min=1,max=30"`
	Age        int    `json:"age" binding:"omitempty,min=1,max=99"`
}
