package load

import "github.com/demoManito/goloader"

func init() {
	goloader.Register(func(loader goloader.ILoader) (err error) {
		load := &Load{
			Id:   1,
			Name: "test_name",
		}
		err = loader.Register("load.load", load)
		return
	})
}

type Load struct {
	Id   int
	Name string
}
