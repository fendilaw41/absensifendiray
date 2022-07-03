package user

type UserRequest struct {
	Name          string `json:"name" form:"name" binding:"required"`
	Email         string `json:"email" form:"email"`
	LastName      string `json:"lastname" form:"lastname"`
	DepartementId int    `json:"departement_id" form:"departement_id"`
	DivisiId      int    `json:"divisi_id" form:"divisi_id"`
	Password      string `json:"password" form:"password"`
}

type UserRequestPUT struct {
	Name          string `json:"name" form:"name" binding:"required"`
	LastName      string `json:"lastname" form:"lastname"`
	DepartementId int    `json:"departement_id" form:"departement_id"`
	DivisiId      int    `json:"divisi_id" form:"divisi_id"`
	Password      string `json:"password" form:"password"`
}
