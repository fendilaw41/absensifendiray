package permission

type PermissionResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SelectPermission struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ResultPermission(u Permission) PermissionResponse {
	return PermissionResponse{
		Id:          u.Id,
		Name:        u.Name,
		Description: u.Description,
	}
}
