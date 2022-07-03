package divisi

type DivisiRequest struct {
	Name string `json:"name" binding:"required"`
}
