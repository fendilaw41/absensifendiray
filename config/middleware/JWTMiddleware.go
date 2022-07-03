package middleware

import (
	"github.com/fendilaw41/absensifendiray/app/src/user"
	"github.com/fendilaw41/absensifendiray/config/action"
	"github.com/fendilaw41/absensifendiray/config/database"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {

	return func(res *gin.Context) {

		clientToken := res.Request.Header.Get("Authorization")
		if clientToken == "" {
			action.NoAccess(res)
			res.Abort()
			return
		}

		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			action.BadRequest(clientToken, res)
			res.Abort()
			return
		}

		jwtWrapper := JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			res.JSON(401, err.Error())
			res.Abort()
			return
		}

		var user user.User
		database, _ := database.DbSetup()
		database.Where("email = ?", claims.Email).Preload("Roles").First(&user)
		var rolesObj string
		for _, v := range user.Roles {
			rolesObj = v.Name
		}
		res.Set("email", claims.Email)
		res.Set("authUser", user)
		res.Set("authId", user.ID)
		res.Set("authName", user.Name)
		res.Set("authRole", rolesObj)
		res.Set("token", clientToken)

		res.Next()

	}
}
