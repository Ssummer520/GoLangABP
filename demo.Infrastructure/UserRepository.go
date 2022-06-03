package repository

import (
	dto "GoLangABP/demo.Core/Dto"
	. "GoLangABP/demo.Core/Model"
	"GoLangABP/demo.Infrastructure/datasource"
	"fmt"
	"k8s.io/klog/v2"
)

// UserRepository  注入数据库
type UserRepository struct {
	SqlDb *datasource.Db `inject:""`
}

// FirstOrDefault FirstOrDefault
func (t *UserRepository) FirstOrDefault(name string, phone string) *UserInfo {
	var user = &UserInfo{}
	err := t.SqlDb.DB().Get(user, "select id, name,phone,age,password from userinfo  where name=? && phone=? limit 1",
		name, phone)
	if err != nil {
		klog.Error(err)
	}
	return user
}

func (t *UserRepository) Add(input dto.UserAddInputDto) dto.UserAddOutPutDto {
	isSuccess := true
	var id int64
	error := ""

	result, err := t.SqlDb.DB().Exec("INSERT INTO userinfo (`name`, `phone`, `age`, `password`) VALUES (?, ?, ?,?)",
		input.Name, input.Phone, input.Age, input.PassWord)
	if err != nil {
		error = err.Error()
		klog.Error(err)
		isSuccess = false
	} else {
		fmt.Println(result)
		id, err = result.LastInsertId()
		if err != nil {
			isSuccess = false
			klog.Error(err)
		}
	}
	ret := &dto.UserAddOutPutDto{
		Success: isSuccess,
		Id:      id,
		Error:   error,
	}
	return *ret
}

func (t *UserRepository) List() []UserInfo {
	var list []UserInfo
	err := t.SqlDb.DB().Select(&list, "select name,age,phone,password from userinfo")
	if err != nil {
		klog.Error(err)
	}
	return list
}
func (t *UserRepository) CheckUserInfo(name string, phone string) bool {
	var user = &UserInfo{}
	err := t.SqlDb.DB().Get(user, "select name,phone,age,password from userinfo  where name=? || phone=? limit 1",
		name, phone)
	if err != nil {
		klog.Error(err)
		if err.Error() == "sql: no rows in result set" {
			return false
		}
		return true
	}
	klog.Info(user)
	if user != nil {
		return true
	}
	return false
}
