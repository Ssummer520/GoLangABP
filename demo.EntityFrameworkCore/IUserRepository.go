package repository

import "GoLangABP/demo.Core/Model"

//IRepository IRepository
type IUserRepository interface {
	FirstOrDefault() *Model.UserInfo
}
