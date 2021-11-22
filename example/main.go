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
	// 2. 初始化
	goloader.LoadingAll(loader)

	h := &Handle{}
	//
	loader.Loading(h)
	l, _ := json.Marshal(h.service.GetLoad())
	fmt.Println(string(l))

	// 直接获取
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
