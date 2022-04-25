package MapperHelper

import . "GoLangABP/demo.Core/Model"

import "github.com/devfeel/mapper"

func init() {
	err := mapper.Register(&UserInfo{})
	if err != nil {
		return
	}
}
func Mapper(fromObj any, toObj any) error {
	err := mapper.Mapper(fromObj, toObj)
	return err
}
func MapperSlice(fromObj any, toObj any) error {
	err := mapper.MapperSlice(fromObj, toObj)
	return err
}
