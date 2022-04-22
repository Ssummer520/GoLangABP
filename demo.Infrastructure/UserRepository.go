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
	err := t.SqlDb.DB().Get(user, "select name,sex,phone,age,password from userinfo limit 1")
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func (t *UserRepository) Add(input dto.UserAddInputDto) dto.UserAddOutPutDto {
	isSuccess := true
	var id int64
	error := ""
	result, err := t.SqlDb.DB().Exec("INSERT INTO userinfo (`name`, `phone`,  `password`) VALUES (?, ?, ?)",
		input.Name, input.Phone, input.PassWord)
	if err != nil {
		isSuccess = false
		fmt.Println("errrrr")
	} else {
		fmt.Println(result)
		id, err = result.LastInsertId()
		if err != nil {
			fmt.Println("errrrr")
			isSuccess = false
			fmt.Println(err)
		}
	}
	fmt.Println(1111)
	ret := &dto.UserAddOutPutDto{
		Success: isSuccess,
		Id:      id,
		Error:   error,
	}
	fmt.Println(1111)
	return *ret
}
