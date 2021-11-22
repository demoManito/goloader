package load

import "github.com/demoManito/goloader"

func init() {
	goloader.Register(func(loader goloader.ILoader) (err error) {
		var dao ILoadDao = new(LoadDao)
		err = loader.Register("load.ILoadDao", dao)
		return
	})
}

type ILoadDao interface {
	GetLoad() *Load
}
type LoadDao struct {
	load *Load `load:"load.load"`
}

func (l *LoadDao) GetLoad() *Load {
	return l.load
}
