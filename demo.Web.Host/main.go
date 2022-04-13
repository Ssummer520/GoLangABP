package main

import (
	repository "GoLangABP/demo.Application/repository"
	service "GoLangABP/demo.Application/start"
	. "GoLangABP/demo.Web.Host/conf"
	. "GoLangABP/demo.Web.Host/controllers"
	_ "GoLangABP/demo.Web.Host/docs" // 千万不要忘了导入把你上一步生成的docs
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
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
// @host            localhost:8889
// @BasePath
/*  @host 124.220.12.138:8888*/
/*  @host localhost:8888*/
func main() {
	/*  @host 124.220.12.138:8888*/
	envPort := ""
	gin.SetMode(gin.DebugMode)
	if gin.Mode() == gin.ReleaseMode {
		envPort = os.Getenv("ASPNETCORE_PORT")
	}

	CarInventoryModel.X = 2
	r := gin.New()
	Configure(r)
	r.GET("/age", GetAgeHandler)
	r.GET("/car", GetCarHandler)
	r.POST("/age", AddAgeHandler)
	r.GET("/rate", GetRateHandler)
	r.GET("/buffer", GetBufferHandler)
	r.GET("/name", index.GetNameHandler)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	if gin.Mode() == gin.ReleaseMode {
		fmt.Println(gin.Mode())
		r.Run(":" + envPort)
	} else {
		r.Run(":8889")
	}
}

func Configure(app *gin.Engine) {
	//controller declare
	var index Index

	//inject declare
	//db := datasource.Db{}
	//redis := cache.Redis{}
	//rabbit := rabbitmq.Mq{}

	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		//&inject.Object{Value: &db},
		//&inject.Object{Value: &redis},
		//&inject.Object{Value: &rabbit},
		&inject.Object{Value: &repository.StartRepo{}},
		&inject.Object{Value: &service.StartService{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//database connect
	//err = db.Connect()
	//if err != nil {
	//	log.Fatal("db fatal:", err)
	//}
	// //redis server connect
	// err = redis.Connect()
	// if err != nil {
	// 	log.Fatal("redis fatal:", err)
	// }
	// // rabbitmq server connect
	// err = rabbit.Connect()
	// if err != nil {
	// 	log.Fatal("RabbitMQ fatal:", err)
	// }

}
