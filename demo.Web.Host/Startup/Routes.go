package Startup

import (
	"GoLangABP/demo.Web.Host/Authentication"
	. "GoLangABP/demo.Web.Host/controllers"

	"github.com/gin-gonic/gin"
)

// Server Index1 路由
type Server struct {
	UserApi      *UserController      `inject:""`
	UserLoginApi *UserLoginController `inject:""`
}

var JwtR = &Authentication.JWTHelper{}
var Api = &Server{}

func ConfigureRoute(r *gin.Engine) {
	//controller declare
	r.POST("/user", Api.UserApi.AddUserNameHandler)
	r.POST("/login", Api.UserLoginApi.LoginHandler)
	r.Use(JwtR.JwtVerify)

}
