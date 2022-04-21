package userService

import (
	dto "GoLangABP/demo.Core/Dto"
	_ "GoLangABP/demo.Core/Model"
	rep "GoLangABP/demo.Infrastructure"
	"fmt"
)
import MapperHelper "GoLangABP/demo.Core/Mapper"

//UserService 注入UserService
type UserService struct {
	userRepository rep.IUserRepository `inject:""`
}

//Login 实现Login方法
func (u *UserService) Login(input dto.UserLoginInputDto) bool {

	var outPut = &dto.UserLoginOutPutDto{}
	fmt.Println(input.PassWord, input.Name)
	userModel := u.userRepository.FirstOrDefault()
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
func (u *UserService) Add(input dto.UserAddInputDto) *dto.UserAddOutPutDto {
	userAddOutPutDto := u.userRepository.Add(input)
	return userAddOutPutDto
}

func (u *UserService) First(input dto.UserLoginInputDto) dto.UserLoginOutPutDto {
	data := u.userRepository.FirstOrDefault()
	var ret = &dto.UserLoginOutPutDto{}
	MapperHelper.Mapper(data, ret)
	return *ret
}
