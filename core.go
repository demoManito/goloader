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
	Register(key interface{}, value interface{}) error
	Replace(key interface{}, value interface{})
	Get(key string) (interface{}, error)
	Remove(key string)
	LoadingAll()
	Loading(v interface{})
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
