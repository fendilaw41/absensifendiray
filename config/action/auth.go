package action

import "github.com/gin-gonic/gin"

func AuthRole(res *gin.Context) (role string) {
	v, ok := res.Get("authRole")
	if ok {
		role = v.(string)
	}
	return
}

func AuthToken(res *gin.Context) (token string) {
	v, ok := res.Get("token")
	if ok {
		token = v.(string)
	}
	return
}

func AuthId(res *gin.Context) (Id int) {
	v, ok := res.Get("authId")
	if ok {
		Id = v.(int)
	}
	return
}

func AuthName(res *gin.Context) (name string) {
	v, ok := res.Get("authName")
	if ok {
		name = v.(string)
	}
	return
}

func AuthUser(res *gin.Context) (user interface{}) {
	v, ok := res.Get("authUser")
	if ok {
		user = v
	}
	return
}

func AuthEmail(res *gin.Context) (email string) {
	v, ok := res.Get("email")
	if ok {
		email = v.(string)
	}
	return
}
