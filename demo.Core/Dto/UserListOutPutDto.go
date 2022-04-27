package dto

type UserListOutPutDto struct {
	Name  string `json:"name" mapper:"Name"`
	Age   int    `json:"age" mapper:"Age"`
	Phone string `json:"phone" mapper:"Phone"`
}
