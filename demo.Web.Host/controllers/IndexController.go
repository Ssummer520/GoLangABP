package controllers

import (
	service "GoLangABP/demo.Application/start"
	. "GoLangABP/demo.Web.Host/Authentication"
	"fmt"
	"github.com/gin-gonic/gin"
)

//Index 注入IStartService
type Index struct {
	Service service.IStartService `inject:""`
	Source  IJWTHelper            `inject:""`
}

// GetNameHandler
// @Summary  获取最新age
// @Tags     Index 依赖注入相关
// @Accept   application/json
// @Produce  application/json
// @Router   /name [get]
// @Param Authorization header string false "Bearer 用户令牌"
// @Success 200 {string} string "{"car":"211221321"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
func (i *Index) GetNameHandler(c *gin.Context) {
	fmt.Println(11111)
	i.Source.JwtVerify(c)
	str := i.Service.Say("11111")
	fmt.Println(str)
	c.JSON(200, gin.H{
		"car": str,
	})
}
