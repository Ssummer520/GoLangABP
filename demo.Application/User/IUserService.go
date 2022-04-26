package userService

import "GoLangABP/demo.Core/Dto"

//IUserService 定义IUserService接口
type IUserService interface {
	//Login 登录
	Login(input dto.UserLoginInputDto) bool
	// Add 新增用户
	Add(input dto.UserAddInputDto) dto.UserAddOutPutDto
	// List 获取用户列表
	List() []dto.UserAddOutPutDto
	//CheckUserInfo 新增用户时判断user信息是否已经存在
	CheckUserInfo(name string, phone string) bool
}
