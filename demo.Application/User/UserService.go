package userService

import (
	dto "GoLangABP/demo.Core/Dto"
	_ "GoLangABP/demo.Core/Model"
	rep "GoLangABP/demo.Infrastructure"
	"k8s.io/klog/v2"
)
import MapperHelper "GoLangABP/demo.Core/Mapper"

//UserService 注入UserService
type UserService struct {
	Rep *rep.UserRepository `inject:""`
}

//Login 实现Login方法
func (u *UserService) Login(input dto.UserLoginInputDto) bool {

	var outPut = &dto.UserLoginOutPutDto{}
	userModel := u.Rep.FirstOrDefault(input.Name, input.PassWord)
	err := MapperHelper.Mapper(userModel, outPut)
	if err != nil {
		klog.Error(err)
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
	klog.Info(retModel)
	return retModel
}
func (u *UserService) CheckUserInfo(name string, phone string) bool {
	isExist := u.Rep.CheckUserInfo(name, phone)
	return isExist
}
