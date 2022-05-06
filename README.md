# GoLangABP

参照ABP(.NET)框架搭建 参考资料https://www.jianshu.com/p/d6c2dc7b5c79
搭建的时候参考很多已有的框

运行起来后 http://localhost:8181/swagger/index.html    项目已经发布在腾讯云服务器 http://124.220.12.138:8889/swagger/index.html

框架 由 Gin + go-JWT +SwaggerUi+facebook依赖注入框架+类似于autoMapper+使用go sqlx做db与实体之间的转换
实现dto 与model之间的数据转换中。
JWT依赖注入
项目初始搭建 正在完善具体细节 完善代码


Swagger还原的时候需要注意 当前的go版本为1.18需要执行 
go install  github.com/swaggo/swag/cmd/swag@latest


![](../../../Desktop/4933701-bb2953451b9b14e7.jpeg)
