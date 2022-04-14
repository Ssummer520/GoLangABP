package controllers

import (
	userService "GoLangABP/demo.Application/user"
	. "GoLangABP/demo.Application/user/dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	User userService.IUserService `inject:""`
}

// LoginHandler
// @Summary  用户登录
// @Tags     UserLogin 登录相关
// @Accept   application/json
// @Produce  application/json
// @Param object body UserLoginInputDto false "请求参数"
// @Router   /login [post]
func (u *UserLogin) LoginHandler(c *gin.Context) {
	var dto UserLoginInputDto
	c.ShouldBind(dto)
	success := u.User.Login(dto)
	fmt.Println(111111111)
	c.JSON(200, gin.H{
		"Success": success,
	})
}
