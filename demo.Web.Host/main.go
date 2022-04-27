package main

import (
	"GoLangABP/demo.Web.Host/Startup"
	_ "GoLangABP/demo.Web.Host/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/gin-gonic/gin"
	"os"
)

// @title           GoLangWebFrameWork
// @version         1.0
// @description     参考ABP 框架搭建
// @termsOfService  http://swagger.io/terms/
// @contact.name    1362
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:8889
// @BasePath

func main() {
	/*  @host localhost:8889
	124.220.12.138:8889

	*/
	envPort := ""
	gin.SetMode(gin.DebugMode)
	if gin.Mode() == gin.ReleaseMode {
		envPort = os.Getenv("ASPNETCORE_PORT")
	}
	r := gin.New()
	Startup.Configure(r)
	if gin.Mode() == gin.ReleaseMode {
		_ = r.Run(":" + envPort)
	} else {
		_ = r.Run(":8889")
	}
}
