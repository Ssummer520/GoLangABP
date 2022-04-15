package main

import (
	. "GoLangABP/demo.Application/user/dto"
	. "GoLangABP/demo.Core/Model"
	"GoLangABP/demo.Web.Host/Startup"
	_ "GoLangABP/demo.Web.Host/docs" // 千万不要忘了导入把你上一步生成的docs
	"fmt"
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"os"
)

// @title           ExamCodeAboutTeslaByGin
// @version         1.0
// @description     ExamCodeAboutTeslaByGin
// @termsOfService  http://swagger.io/terms/
// @contact.name    1362
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8181
// @BasePath
/*  @host 124.220.12.138:8888*/
/*  @host localhost:8888*/
func init() {
	err := mapper.Register(&UserInfo{})
	if err != nil {
		return
	}
}
func main() {
	/*  @host 124.220.12.138:8888*/
	envPort := ""
	gin.SetMode(gin.DebugMode)
	if gin.Mode() == gin.ReleaseMode {
		envPort = os.Getenv("ASPNETCORE_PORT")
	}
	var userInfo = &UserInfo{}
	userInfo.Phone = "1362246612"
	userInfo.Age = 1
	userInfo.Sex = 1
	userInfo.Name = "aa"
	var outPut = &UserLoginOutPutDto{}
	err := mapper.Mapper(userInfo, outPut)
	if err != nil {
		return
	}
	r := gin.New()
	Startup.Configure(r)
	if gin.Mode() == gin.ReleaseMode {
		fmt.Println(gin.Mode())
		_ = r.Run(":" + envPort)
	} else {
		_ = r.Run(":8181")
	}
}
