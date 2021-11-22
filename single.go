package goloader

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

var _ ILoader = new(singleLoader)

func NewSingleLoader() ILoader {
	return &singleLoader{
		objs: make(map[interface{}]reflect.Value),
	}
}

type singleLoader struct {
	objs map[interface{}]reflect.Value
}

func (s *singleLoader) Register(key interface{}, value interface{}) error {
	_, ok := s.objs[key]
	if ok {
		return errors.New(fmt.Sprintf("key duplicate: %v", key))
	}
	s.objs[key] = reflect.ValueOf(value)
	return nil
}

func (s *singleLoader) Replace(key interface{}, value interface{}) {
	s.objs[key] = reflect.ValueOf(value)
}

func (s *singleLoader) Get(key string) (interface{}, error) {
	v, ok := s.objs[key]
	if ok {
		return v.Interface(), nil
	}
	return nil, ErrNotFound
}

func (s *singleLoader) Remove(key string) {
	delete(s.objs, key)
}

func (s *singleLoader) Clear() {
	s.objs = make(map[interface{}]reflect.Value)
}

func (s *singleLoader) LoadingAll() {
	for _, v := range s.objs {
		s.Loading(v)
	}
}

// 该方法会一层一层递归结构体，将结构体所有的列全部赋值（前提是load标签数据已加载在内存中）
func (s *singleLoader) Loading(v interface{}) {
	var value reflect.Value
	var ok bool
	if value, ok = v.(reflect.Value); !ok {
		value = reflect.ValueOf(v)
	}
loop:
	for {
		switch value.Kind() {
		case reflect.Ptr:
			value = value.Elem()
		case reflect.Interface:
			value = value.Elem()
		default:
			break loop
		}
	}

	if value.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < value.NumField(); i++ {
		name := value.Type().Field(i).Tag.Get(LoaderTag)
		temp, ok := s.objs[name]
		if ok {
			field := value.Field(i)
			if field.CanSet() {
				field.Set(temp)
			} else {
				field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
				field.Set(temp)
			}
		}
	}
}
