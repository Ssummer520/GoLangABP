package Model

type UserInfo struct {
	//Name  string `gorm:"name"`
	//Sex   int    `gorm:"sex"`
	//Age   int    `gorm:"age"`
	//Phone string `gorm:"phone"`
	Id   string
	Name string
	//Sex  int
	Age      int
	Phone    string
	PassWord string
}
