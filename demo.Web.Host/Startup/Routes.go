package Startup

import (
	"GoLangABP/demo.Web.Host/Authentication"
	. "GoLangABP/demo.Web.Host/controllers"
	"github.com/gin-gonic/gin"
)

//路由
var IndexC Index
var UserC UserLogin

func ConfigureRoute(r *gin.Engine) {
	//controller declare

	r.Use(Authentication.JwtVerify)
}
