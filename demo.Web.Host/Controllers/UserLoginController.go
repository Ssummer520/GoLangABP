package Controllers

import (
	userService "GoLangABP/demo.Application/User"
	. "GoLangABP/demo.Core/Dto"
	"GoLangABP/demo.Core/Model"
	. "GoLangABP/demo.Web.Host/Authentication"
	jwt "GoLangABP/demo.Web.Host/Authentication"
	"github.com/gin-gonic/gin"
)

type UserLoginController struct {
	*userService.UserService `inject:""`
	Jwt                      IJWTHelper `inject:""`
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
func (u *UserLoginController) LoginHandler(c *gin.Context) {
	var dto UserLoginInputDto
	c.ShouldBind(&dto)
	success, userId := u.Login(dto)
	message := ""
	token := ""
	if success {
		claims := &jwt.UserClaims{}
		claims.ID = userId
		//claims.Phone = ""
		//claims.Name = ""
		token = u.Jwt.GenerateToken(claims)
	}

	retObject := Model.RetObject{
		Success: success,
		Data:    token,
		Message: message,
	}
	c.JSON(200, retObject)
}
