package userService

import (
	"GoLangABP/demo.Application/user/dto"
	"fmt"
)

//UserService 注入UserService
type UserService struct {
}

//Login 实现Login方法
func (u *UserService) Login(input dto.UserLoginInputDto) bool {
	fmt.Println(input)

	if input.PassWord == "" || input.Name == "" {
		return false
	}
	return true
}
