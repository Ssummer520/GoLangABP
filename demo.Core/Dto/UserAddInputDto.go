package dto

// UserAddInputDto 新建用户
type UserAddInputDto struct {
	Name     string `json:"name"`     //姓名
	Sex      int    `json:"sex"`      //性别
	Age      int    `json:"age"`      //年龄
	Phone    string `json:"phone"`    //电话
	PassWord string `json:"password"` //密码
}
