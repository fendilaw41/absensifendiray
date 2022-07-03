package controllers

import (
	"log"
	"net/http"

	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"
	"github.com/fendilaw41/absensifendiray/config/middleware"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func constructLogin() *repository {
	db, _ := database.DbSetup()
	return &repository{db}
}

// Signup creates a user in db
func Signup(res *gin.Context) {
	var userModel user.User
	var req user.UserRequest
	err := res.ShouldBindJSON(&req)
	if err != nil {
		res.JSON(400, gin.H{
			"message": "Anda Belum Memasukan Apapun",
		})
		return
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), 14)
	if err != nil {
		res.JSON(500, gin.H{
			"message": "Gagal Men-Generate Password",
		})
		return
	}
	userModel.Password = string(bytes)
	userModel.Email = req.Email
	userModel.Name = req.Name
	userModel.FirstName = req.Name
	userModel.LastName = req.LastName
	userModel.Fullname = req.Name + " " + req.LastName
	userModel.DepartementId = req.DepartementId
	userModel.DivisiId = req.DivisiId
	err = constructLogin().db.Create(&userModel).Error
	if err != nil {
		action.BadRequest(err.Error(), res)
		return
	}
	var roles []user.RolesResponse
	for _, v := range userModel.Roles {
		newRespon := user.RolesResponse{
			Name: v.Name,
		}
		roles = append(roles, newRespon)
	}
	registerResp := resultResponseRegister(userModel, roles)
	action.AccepData("Sukses Register", registerResp, res)
}

// Login logs users in
func Login(c *gin.Context) {
	var payload user.LoginPayload
	var userModel user.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Anda Belum Memasukan Apapun",
		})
		c.Abort()
		return
	}
	result := constructLogin().db.Where("email = ?", payload.Email).Preload("Roles.Permission").First(&userModel)
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"message": "Email Anda Salah",
		})
		c.Abort()
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(payload.Password))
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"message": "Password Anda Salah",
		})
		c.Abort()
		return
	}
	jwtWrapper := middleware.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}
	signedToken, err := jwtWrapper.GenerateToken(userModel.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Token Salah Generate",
		})
		c.Abort()
		return
	}
	var rolesArray []user.RolesResponse
	for _, v := range userModel.Roles {
		newRoles := user.RolesResponse{
			Name: v.Name,
		}
		rolesArray = append(rolesArray, newRoles)
	}
	tokenResponse := user.LoginResponse{
		Token: signedToken,
		Roles: rolesArray,
	}

	action.AccepData("Berhasil Mendapatkan Token", tokenResponse, c)
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	a := action.Default(c)
	action.Logout(session, a.User)

	c.Redirect(http.StatusMovedPermanently, "/")

	action.AccepMsg("berhasil Logout", c)
}

func Profile(res *gin.Context) {
	var userModel user.User
	authRole := action.AuthRole(res)
	authEmail := action.AuthEmail(res) // from the authorization middleware
	result := constructLogin().db.Where("email = ?", authEmail).Preload("Roles.Permission").First(&userModel)
	if result.Error == gorm.ErrRecordNotFound {
		action.NotFound(res)
		res.Abort()
		return
	}
	if result.Error != nil {
		action.Abort(res)
		res.Abort()
		return
	}
	var roles []user.RolesResponse
	for _, v := range userModel.Roles {
		var permission []user.PermissionResponse
		for _, x := range v.Permission {
			newPermission := user.PermissionResponse{
				Name: x.Name,
			}
			permission = append(permission, newPermission)
		}
		newRoles := user.RolesResponse{
			Name:       v.Name,
			Permission: permission,
		}
		roles = append(roles, newRoles)
	}
	resultRespon := resultResponseRegister(userModel, roles)
	if authRole == "superadmin" {
		action.AccepData("Sukses Get Profile", resultRespon, res)
	} else {
		action.NoAccess(res)
		return
	}
}

func resultResponseRegister(b user.User, roles interface{}) user.UserResponse {
	return user.UserResponse{
		ID:    b.ID,
		Name:  b.Name,
		Email: b.Email,
		Roles: roles,
	}
}
