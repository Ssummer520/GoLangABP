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
	User *userService.UserService `inject:""`
	Jwt  IJWTHelper               `inject:""`
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
	success := u.User.Login(dto)
	message := ""
	token := ""
	if success {
		claims := &jwt.UserClaims{}
		claims.ID = "1"
		claims.Phone = "17343016071"
		claims.Name = "17343016071"
		token = u.Jwt.GenerateToken(claims)
	}

	object := Model.RetObject{
		Success: success,
		Data:    token,
		Message: message,
	}
	c.JSON(200, object)
}
