package dto

type UserAddOutPutDto struct {
	Success bool   `json:"success"` //姓名
	Error   string `json:"error"`   //性别
	Id      int64  `json:"userid"`  //年龄
}
