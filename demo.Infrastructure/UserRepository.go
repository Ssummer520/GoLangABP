package repository

import (
	dto "GoLangABP/demo.Core/Dto"
	. "GoLangABP/demo.Core/Model"
	"GoLangABP/demo.Infrastructure/datasource"
	"fmt"
)

// UserRepository  注入数据库
type UserRepository struct {
	Source datasource.IDb `inject:""`
}

// FirstOrDefault FirstOrDefault
func (t *UserRepository) FirstOrDefault() *UserInfo {
	var user = &UserInfo{}
	err := t.Source.DB().Get(user, "select name,sex,phone,age,password from userinfo limit 1")
	if err != nil {
		fmt.Println(err)
	}
	return user
}
func (t *UserRepository) Add(input dto.UserAddInputDto) bool {
	isSuccess := true
	result, err := t.Source.DB().Exec("INSERT INTO userinfo (`name`, `phone`,  `password`) VALUES (?, ?, ?)",
		input.Name, input.Phone, input.PassWord)
	if err != nil {
		isSuccess = false
		fmt.Println(err)
	}
	result.LastInsertId()
	return isSuccess
}
