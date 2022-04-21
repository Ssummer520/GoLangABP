package Startup

import (
	service "GoLangABP/demo.Application/start"
	userService "GoLangABP/demo.Application/user"
	rep "GoLangABP/demo.Infrastructure"
	"GoLangABP/demo.Infrastructure/datasource"
	. "GoLangABP/demo.Web.Host/Authentication"
	. "GoLangABP/demo.Web.Host/controllers"
	_ "GoLangABP/demo.Web.Host/docs" // 千万不要忘了导入把你上一步生成的docs
	"github.com/facebookgo/inject"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
)

func Configure(r *gin.Engine) {

	//controller declare
	var index Index
	var userLogin UserLogin
	indexR = &index
	userLoginR = &userLogin
	//inject declare
	db := datasource.Db{}
	//redis := cache.Redis{}
	//rabbit := rabbitmq.Mq{}

	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		&inject.Object{Value: &userLogin},
		&inject.Object{Value: &db},
		//&inject.Object{Value: &repository.StartRepo{}},
		&inject.Object{Value: &service.StartService{}},
		&inject.Object{Value: &rep.UserRepository{}},
		&inject.Object{Value: &userService.UserService{}},
		&inject.Object{Value: &JWTHelper{}},

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
