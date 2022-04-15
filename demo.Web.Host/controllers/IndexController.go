package controllers

import (
	service "GoLangABP/demo.Application/start"
	"fmt"
	"github.com/gin-gonic/gin"
)

//Index 注入IStartService
type Index struct {
	Service service.IStartService `inject:""`
}

// GetNameHandler
// @Summary  获取最新age
// @Tags     Index 依赖注入相关
// @Accept   application/json
// @Produce  application/json
// @Router   /name [get]
func (i *Index) GetNameHandler(c *gin.Context) {

	fmt.Println(11111)
	str := i.Service.Say("11111")
	fmt.Println(str)
	c.JSON(200, gin.H{
		"car": str,
	})
}
