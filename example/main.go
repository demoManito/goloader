package main

import (
	"encoding/json"
	"fmt"

	"github.com/demoManito/goloader"
	"github.com/demoManito/goloader/example/load"
)

func main() {
	// 1. 获取单例对象
	loader := goloader.DefaultLoader()
	// 2. 将数据加载进入 loader
	goloader.LoadingAll(loader)

	// 基本使用方法
	h := &Handle{}
	// 3-1. 自动装载
	// 将 loader 池中的值，赋值给变量 h
	loader.Loading(h)
	l, _ := json.Marshal(h.service.GetLoad())
	fmt.Println(string(l))

	// 基本使用方法
	// 3-2. 直接获取
	inter, err := loader.Get("load.load")
	if err != nil {
		if err == goloader.ErrNotFound {
			panic("Not found")
		}
		panic(err)
	}
	fmt.Println(inter)
}

type Handle struct {
	service load.ILoadService `load:"load.ILoadService"`
}

func (h *Handle) GetUser() *load.Load {
	s := h.service
	return s.GetLoad()
}
