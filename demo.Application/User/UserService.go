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
	Rep *rep.UserRepository `inject:""`
}

//Login 实现Login方法
func (u *UserService) Login(input dto.UserLoginInputDto) bool {

	var outPut = &dto.UserLoginOutPutDto{}
	userModel := u.Rep.FirstOrDefault()
	err := MapperHelper.Mapper(userModel, outPut)
	if err != nil {
		fmt.Println(err)
	}
	if input.PassWord != outPut.PassWord || input.Name != outPut.Name {
		return false
	}
	return true
}
func (u *UserService) Add(input dto.UserAddInputDto) dto.UserAddOutPutDto {
	userAddOutPutDto := u.Rep.Add(input)
	return userAddOutPutDto
}
func (u *UserService) List() []dto.UserListOutPutDto {
	userInfo := u.Rep.List()
	var retModel []dto.UserListOutPutDto
	MapperHelper.MapperSlice(&userInfo, &retModel)
	fmt.Println(retModel)
	return retModel
}
