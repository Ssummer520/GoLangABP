package Startup

import (
	"GoLangABP/demo.Web.Host/Authentication"
	. "GoLangABP/demo.Web.Host/controllers"
	"github.com/gin-gonic/gin"
)

// Index1 路由
var indexR = &Index{}
var userLoginR = &UserLogin{}

func ConfigureRoute(r *gin.Engine) {
	//controller declare
	r.POST("/login", userLoginR.LoginHandler)
	r.Use(Authentication.JwtVerify)
	r.GET("/name", indexR.GetNameHandler)
}
