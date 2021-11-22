package goloader

import "errors"

type LoadFunc func(ILoader) error

const LoaderTag = "load"

var (
	ErrNotFound = errors.New("key not found")
	shared      = NewSingleLoader()
	funcs       = make([]LoadFunc, 0, 16)
)

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

func Register(f LoadFunc) {
	funcs = append(funcs, f)
}

func LoadingAll(loader ILoader) (err error) {
	for _, f := range funcs {
		err = f(loader)
		if err != nil {
			return err
		}
	}
	loader.LoadingAll()
	return
}
