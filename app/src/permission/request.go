package permission

type PermissionRequest struct {
	Name        string `json:"name" form:"name" binding:"required"`
	Description string `json:"description" form:"description"`
}

type PermissionUpdateRequest struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}
