package Startup

import (
	"GoLangABP/demo.Web.Host/Authentication"
	. "GoLangABP/demo.Web.Host/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoute(r *gin.Engine) {
	//controller declare
	var index Index
	var userLogin UserLogin
	r.POST("/login", userLogin.LoginHandler)
	r.GET("/name", index.GetNameHandler)
	r.Use(Authentication.JwtVerify)
}
