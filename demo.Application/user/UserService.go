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
	IRepository rep.IRepository[UserInfo] `inject:""`
}

//Login 实现Login方法
func (u *UserService) Login(input dto.UserLoginInputDto) bool {
	fmt.Println(input)
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
	if input.PassWord == "" || input.Name == "" {
		return false
	}
	return true
}
func (u *UserService) First(input dto.UserLoginInputDto) dto.UserLoginOutPutDto {
	data := u.IRepository.FirstOrDefault()
	var ret = &dto.UserLoginOutPutDto{}
	MapperHelper.Mapper(data, ret)
	return *ret
}
