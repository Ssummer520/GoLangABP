package controllers

import (
	userService "GoLangABP/demo.Application/user"
	. "GoLangABP/demo.Application/user/dto"
	. "GoLangABP/demo.Web.Host/Authentication"
	jwt "GoLangABP/demo.Web.Host/Authentication"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	User   userService.IUserService `inject:""`
	Source IJWTHelper               `inject:""`
}

// LoginHandler
// @Summary  用户登录
// @Tags     UserLogin 登录相关
// @Accept   application/json
// @Produce  application/json
// @Param object body UserLoginInputDto false "请求参数"
// @Router   /login [post]
// @Success 200 {string} string "{"Success": true,"Token":"5545"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
func (u *UserLogin) LoginHandler(c *gin.Context) {
	var dto UserLoginInputDto
	c.ShouldBind(&dto)
	success := u.User.Login(dto)
	token := ""
	if success {
		claims := &jwt.UserClaims{}
		claims.ID = "1"
		claims.Phone = "17343016071"
		claims.Name = "17343016071"
		token = u.Source.GenerateToken(claims)
	}

	c.JSON(200, gin.H{
		"Success": success,
		"Token":   token,
	})
}
