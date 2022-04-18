package repository

import (
	. "GoLangABP/demo.Core/Model"
	"GoLangABP/demo.EntityFrameworkCore/datasource"
	"fmt"
)

// UserRepository  注入数据库
type UserRepository struct {
	Source datasource.IDb `inject:""`
}

// FirstOrDefault FirstOrDefault
func (t *UserRepository) FirstOrDefault() *UserInfo {
	var user = &UserInfo{}
	fmt.Println(22222222)
	t.Source.DB().Raw("select * from userinfo limit 1").Scan(user)
	fmt.Println(user)
	return user
}
