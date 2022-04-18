package repository

import (
	"github.com/bingjian-zhu/gin-inject/common/datasource"
)

//StartRepo 注入数据库
type Repository[T any] struct {
	Source datasource.IDb `inject:""`
}

//Speak 实现Speak方法
func (t *Repository[T]) FirstOrDefault() any {
	t.Source.DB().First(&t)
	return t
}
