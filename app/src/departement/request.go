package departement

type DepartementRequest struct {
	Name string `json:"name" binding:"required"`
}
