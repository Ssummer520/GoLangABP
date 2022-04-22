package Startup

import (
	repository "GoLangABP/demo.Infrastructure"
	"GoLangABP/demo.Infrastructure/datasource"
	. "GoLangABP/demo.Web.Host/Authentication"
	_ "GoLangABP/demo.Web.Host/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/facebookgo/inject"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

func Configure(r *gin.Engine) {

	db := datasource.Db{}
	//redis := cache.Redis{}
	//rabbit := rabbitmq.Mq{}
	var server = &Server{}
	var jwtHelper = &JWTHelper{}
	Api = server
	JwtR = jwtHelper

	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: server},
		&inject.Object{Value: &repository.UserRepository{}},
		&inject.Object{Value: &db},

		&inject.Object{Value: jwtHelper},

		//&inject.Object{Value: &redis},
		//&inject.Object{Value: &rabbit},
	)

	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}

	//database connect
	err = db.Connect()
	if err != nil {
		log.Fatal("db fatal:", err)
	}

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
	ConfigureRoute(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(cors.Default())
	r.Use(gin.Recovery())
	return
}
