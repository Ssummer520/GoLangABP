package userService

import (
	"GoLangABP/demo.Application/user/dto"
	_ "GoLangABP/demo.Core/Model"
	rep "GoLangABP/demo.EntityFrameworkCore"
	"fmt"
)
import MapperHelper "GoLangABP/demo.Core/Mapper"

//UserService 注入UserService
type UserService struct {
	Source rep.IUserRepository `inject:""`
}

//Login 实现Login方法
func (u *UserService) Login(input dto.UserLoginInputDto) bool {

	var outPut = &dto.UserLoginOutPutDto{}

	fmt.Println(input.PassWord, input.Name)
	userModel := u.Source.FirstOrDefault()

	err := MapperHelper.Mapper(userModel, outPut)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input.Name, outPut.Name, input.PassWord, outPut.PassWord)
	if input.PassWord != outPut.PassWord || input.Name != outPut.Name {
		return false
	}
	return true
}

func (u *UserService) First(input dto.UserLoginInputDto) dto.UserLoginOutPutDto {
	data := u.Source.FirstOrDefault()
	var ret = &dto.UserLoginOutPutDto{}
	MapperHelper.Mapper(data, ret)
	return *ret
}
