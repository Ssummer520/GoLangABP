package Controllers

import (
	service "GoLangABP/demo.Application/User"
	. "GoLangABP/demo.Core/Dto"
	"fmt"
	"github.com/gin-gonic/gin"
)

// User  注入IStartService
type UserController struct {
	Service *service.UserService `inject:""`
}

// AddUserNameHandler
// @Summary  新增用户信息
// @Tags     User 用户相关
// @Accept   application/json
// @Produce  application/json
// @Router   /User [post]
// @Param object body UserAddInputDto false "请求参数"
func (u *UserController) AddUserNameHandler(c *gin.Context) {
	var addModel = UserAddInputDto{}
	c.ShouldBind(&addModel)
	ret := u.Service.Add(addModel)
	fmt.Println(ret)
	c.JSON(200, gin.H{
		"success": ret.Success,
		"error":   ret.Error,
	})
}
