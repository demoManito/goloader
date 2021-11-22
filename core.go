package goloader

import "errors"

const LoaderTag = "load"

var (
	ErrNotFound = errors.New("key not found")
	shared = NewSingleLoader()
)

type LoadFunc func(ILoader)error
func DefaultLoader() ILoader {
	return shared
}

type ILoader interface {
	// 注册
	Register(key interface{}, value interface{}) error
	// 替换
	Replace(key interface{}, value interface{})
	// 获取
	Get(key string) (interface{}, error)
	// 删除
	Remove(key string)
	// 加载所有注册变量
	LoadingAll()
	// 加载某结构体下变量
	Loading(v interface{})
	// 删除 loader 中所有数据
	Clear()
}

