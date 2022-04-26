package repository

import (
	. "GoLangABP/demo.Core/Dto"
	"GoLangABP/demo.Core/Model"
)

type IUserRepository interface {
	FirstOrDefault(name string, phone string) *Model.UserInfo
	Add(input UserAddInputDto) UserAddOutPutDto
	// List 获取user列表信息
	List() []Model.UserInfo
	//CheckUserInfo 新增用户时判断user信息是否已经存在
	CheckUserInfo(name string, phone string) bool
}
