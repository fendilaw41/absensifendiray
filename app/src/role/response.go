package role

import "github.com/fendilaw41/absensifendiray/app/src/permission"

type RolesResponse struct {
	Id         int         `json:"id"`
	Name       string      `json:"role_name"`
	Permission interface{} `json:"permission"`
}

type PermissionResponse struct {
	Name string `json:"name"`
}

func ResultRole(u Role, permission interface{}) RolesResponse {
	return RolesResponse{
		Id:         int(u.Id),
		Name:       u.Name,
		Permission: permission,
	}
}

func ResultPermission(u permission.Permission) PermissionResponse {
	return PermissionResponse{
		Name: u.Name,
	}
}
