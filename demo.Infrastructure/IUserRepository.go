package repository

import (
	. "GoLangABP/demo.Core/Dto"
	"GoLangABP/demo.Core/Model"
)

type IUserRepository interface {
	FirstOrDefault() *Model.UserInfo
	Add(input UserAddInputDto) *UserAddOutPutDto
}
