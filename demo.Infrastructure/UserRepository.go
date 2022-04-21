package repository

import (
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
