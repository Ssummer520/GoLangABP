package repository

import (
	dto "GoLangABP/demo.Core/Dto"
	. "GoLangABP/demo.Core/Model"
	"GoLangABP/demo.Infrastructure/datasource"
	"fmt"
)

// UserRepository  注入数据库
type UserRepository struct {
	SqlDb *datasource.Db `inject:""`
}

// FirstOrDefault FirstOrDefault
func (t *UserRepository) FirstOrDefault() *UserInfo {
	var user = &UserInfo{}
	err := t.SqlDb.DB().Get(user, "select name,phone,age,password from userinfo limit 1")
	if err != nil {
		fmt.Println(err)
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
		isSuccess = false
	} else {
		fmt.Println(result)
		id, err = result.LastInsertId()
		if err != nil {
			isSuccess = false
			error = err.Error()
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
		fmt.Println(11111)
		fmt.Println(err)
	}
	fmt.Println(list)
	return list
}
