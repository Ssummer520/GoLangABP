package userService

import "GoLangABP/demo.Core/Dto"

//IUserService 定义IUserService接口
type IUserService interface {
	//Login 登录
	Login(input dto.UserLoginInputDto) bool
}
