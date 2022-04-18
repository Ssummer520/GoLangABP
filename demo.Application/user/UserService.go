package userService

import (
	"GoLangABP/demo.Application/user/dto"
	. "GoLangABP/demo.Core/Model"
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

	var userInfo = &UserInfo{}
	userInfo.Phone = "1362246612"
	userInfo.Age = 1
	userInfo.Sex = 1
	userInfo.Name = "aa"
	var outPut = &dto.UserLoginOutPutDto{}
	err := MapperHelper.Mapper(userInfo, outPut)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	data := u.Source.FirstOrDefault()
	fmt.Println(data)
	if input.PassWord == "" || input.Name == "" {
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
