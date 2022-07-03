package controllers

import (
	"fmt"
	"strconv"

	"github.com/fendilaw41/absensifendiray/app/src/permission"
	"github.com/fendilaw41/absensifendiray/app/src/role"
	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userController struct {
	userService user.Service
}

type roleController struct {
	roleService role.Service
}

type permissionController struct {
	permissionService permission.Service
}

func constUser() *userController {
	db, _ := database.DbSetup()
	userRepository := user.UserRepository(db)
	userService := user.UserService(userRepository)
	return &userController{userService}
}

func constRole() *roleController {
	db, _ := database.DbSetup()
	roleRepository := role.RoleRepository(db)
	roleService := role.RoleService(roleRepository)
	return &roleController{roleService}
}

func constPermission() *permissionController {
	db, _ := database.DbSetup()
	permissionRepository := permission.PermissionRepository(db)
	permissionService := permission.PermissionService(permissionRepository)
	return &permissionController{permissionService}
}

func GetAllUserPaginate(res *gin.Context) {
	users, page, pageSize, count, err := constUser().userService.FindAllPaginate(res)
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var userModel []user.UserResponse
	for _, u := range users {
		var roles []user.RolesResponse
		for _, v := range u.Roles {
			var permission []user.PermissionResponse
			for _, x := range v.Permission {
				newPermission := user.ResultPermission(x)
				permission = append(permission, newPermission)
			}
			newRoles := user.ResultRole(v, permission)
			roles = append(roles, newRoles)
		}
		newUser := user.ResultUser(u, roles)
		userModel = append(userModel, newUser)
	}
	if userModel == nil {
		action.NotFound(res)
		return
	}
	action.AcceptPaginate(res.Request, res, action.All("User"), userModel, len(userModel), page, pageSize, count)
}

func GetDetailUser(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constUser().userService.FindId(int(id))
	if u.ID == 0 {
		action.NotFound(res)
		return
	}
	if err != nil {
		action.BadRequest(err, res)
	}
	var roles []user.RolesResponse
	for _, v := range u.Roles {
		var permission []user.PermissionResponse
		for _, x := range v.Permission {
			newPermission := user.ResultPermission(x)
			permission = append(permission, newPermission)
		}
		newRoles := user.ResultRole(v, permission)
		roles = append(roles, newRoles)
	}

	userResponse := user.ResultUser(u, roles)
	action.AccepData(action.Detail("User"), userResponse, res)
}

func StoreUser(res *gin.Context) {
	var userReq user.UserRequest
	err := res.ShouldBind(&userReq)

	if err != nil {
		fmt.Println(err)
		action.BadRequest(err, res)
		return
	}

	errService := constUser().userService.Save(userReq, res)
	if errService != nil {
		action.BadRequest(err, res)
		return
	}

	action.AccepMsg(action.Save("User"), res)
}

func UpdateUser(res *gin.Context) {
	var userReq user.UserRequestPUT
	err := res.ShouldBind(&userReq)
	if err != nil {
		fmt.Println(err)
		action.BadRequest(err, res)
		return
	}
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	errService := constUser().userService.UpdateUserService(id, userReq, res)
	if errService != nil {
		action.BadRequest(err, res)
	}

	action.AccepMsg(action.Update("User"), res)
}

func DeleteUser(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	err := constUser().userService.SrcDeleteUser(int(id), res)
	if err != nil {
		action.BadRequest(err, res)
	}

	action.AccepMsg(action.Delete("User"), res)
}

// TODO : SECTION ROLE

func GetAllRolePaginate(res *gin.Context) {
	roles, err := constRole().roleService.FindAll()
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var rolesModel []role.RolesResponse
	for _, u := range roles {
		var permission []role.PermissionResponse
		for _, x := range u.Permission {
			newPermission := role.ResultPermission(x)
			permission = append(permission, newPermission)
		}
		newRole := role.ResultRole(u, permission)
		rolesModel = append(rolesModel, newRole)
	}
	if rolesModel == nil {
		action.NotFound(res)
		return
	}
	action.AccepCount(action.All("Role"), len(rolesModel), rolesModel, res)
}

func GetDetailRole(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constRole().roleService.FindId(int(id))
	if u.ID == 0 {
		action.NotFound(res)
		return
	}
	if err != nil {
		action.BadRequest(err, res)
	}

	var permission []role.PermissionResponse
	for _, x := range u.Permission {
		newPermission := role.ResultPermission(x)
		permission = append(permission, newPermission)
	}
	newRoles := role.ResultRole(u, permission)

	action.AccepData(action.Detail("Role"), newRoles, res)
}

func StoreRole(res *gin.Context) {
	var req role.RoleRequest
	err := res.ShouldBind(&req)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}

	errService := constRole().roleService.SaveRole(req, res)
	if errService != nil {
		action.BadRequest(errService, res)
		return
	}

	action.AccepMsg(action.Save("Role"), res)
}

func UpdateRole(res *gin.Context) {
	var req role.RoleUpdateRequest
	err := res.ShouldBind(&req)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error: %s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	errService := constRole().roleService.UpdateRoleService(id, req, res)
	if errService != nil {
		action.BadRequest(err, res)
	}
	action.AccepMsg(action.Update("Role"), res)
}

func DeleteRole(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	err := constRole().roleService.SrcDeleteRole(int(id))
	if err != nil {
		action.BadRequest(err, res)
	}

	action.AccepMsg(action.Delete("Role"), res)
}

// TODO : SECTION PERMISSION

func GetAllPermissionPaginate(res *gin.Context) {
	p, page, pageSize, count, err := constPermission().permissionService.FindAllPaginate(res)
	if err != nil {
		action.BadRequest(err, res)
		return
	}
	var permissionModel []permission.PermissionResponse
	for _, u := range p {
		newPermission := permission.ResultPermission(u)
		permissionModel = append(permissionModel, newPermission)
	}
	if permissionModel == nil {
		action.NotFound(res)
		return
	}
	action.AcceptPaginate(res.Request, res, action.All("Permission"), permissionModel, len(permissionModel), page, pageSize, count)
}

func GetDetailPermission(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	u, err := constPermission().permissionService.FindId(int(id))
	newPermission := permission.ResultPermission(u)
	if u.Id == 0 {
		action.NotFound(res)
		return
	}
	if err != nil {
		action.BadRequest(err, res)
	}

	action.AccepData(action.Detail("Permission"), newPermission, res)
}

func StorePermission(res *gin.Context) {
	var req permission.PermissionRequest
	err := res.ShouldBind(&req)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}

	errService := constPermission().permissionService.Save(req, res)
	if errService != nil {
		action.BadRequest(errService, res)
		return
	}

	action.AccepMsg(action.Save("Permission"), res)
}

func UpdatePermission(res *gin.Context) {
	var req permission.PermissionUpdateRequest
	err := res.ShouldBind(&req)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error: %s, %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		action.BadRequest(errorMessages, res)
		return
	}
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	errService := constPermission().permissionService.Update(id, req, res)
	if errService != nil {
		action.BadRequest(err, res)
	}
	action.AccepMsg(action.Update("Permission"), res)
}

func DeletePermission(res *gin.Context) {
	idString := res.Param("id")
	id, _ := strconv.Atoi(idString)
	err := constPermission().permissionService.Delete(int(id))
	if err != nil {
		action.BadRequest(err, res)
	}

	action.AccepMsg(action.Delete("Permission"), res)
}
