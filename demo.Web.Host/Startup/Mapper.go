package Startup

import (
	. "GoLangABP/demo.Core/Model"
	"github.com/devfeel/mapper"
)

func init() {
	err := mapper.Register(&UserInfo{})
	if err != nil {
		return
	}
}
