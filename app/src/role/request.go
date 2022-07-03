package role

type RoleRequest struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type RoleUpdateRequest struct {
	Name string `json:"name" form:"name"`
}
