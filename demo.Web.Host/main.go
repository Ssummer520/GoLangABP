package main

import (
	. "GoLangABP/demo.Web.Host/conf"
	. "GoLangABP/demo.Web.Host/controllers"
	_ "GoLangABP/demo.Web.Host/docs" // 千万不要忘了导入把你上一步生成的docs
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)

// @title ExamCodeAboutTeslaByGin
// @version 1.0
// @description ExamCodeAboutTeslaByGin
// @termsOfService http://swagger.io/terms/
// @contact.name 1362
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8888
// @BasePath
/* @host 124.220.12.138:8888*/
/* @host localhost:8888*/
func main() {
	/* @host 124.220.12.138:8888*/
	envPort := ""
	gin.SetMode(gin.ReleaseMode)
	if gin.Mode() == gin.ReleaseMode {
		envPort = os.Getenv("ASPNETCORE_PORT")
	}
	fmt.Println(gin.Mode())
	CarInventoryModel.X = 2

	r := gin.New()
	r.GET("/age", GetAgeHandler)
	r.GET("/car", GetCarHandler)
	r.POST("/age", AddAgeHandler)
	r.GET("/rate", GetRateHandler)
	r.GET("/buffer", GetBufferHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	if gin.Mode() == gin.ReleaseMode {
		fmt.Println(gin.Mode())
		r.Run(":" + envPort)
	} else {
		r.Run(":8888")
	}
}

