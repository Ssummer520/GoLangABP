package Model

import . "GoLangABP/demo.Core/Model"
import "github.com/devfeel/mapper"

func init() {
	err := mapper.Register(&UserInfo{})
	if err != nil {
		return
	}
}

//var userInfo = &UserInfo{}
//userInfo.Phone = "1362246612"
//userInfo.Age = 1
//userInfo.Sex = 1
//userInfo.Name = "aa"
//var outPut = &UserLoginOutPutDto{}
//err := mapper.Mapper(userInfo, outPut)
//if err != nil {
//return
//}
//fmt.Println(outPut)