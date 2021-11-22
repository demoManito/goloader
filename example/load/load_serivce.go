package load

import (
	"github.com/demoManito/goloader"
)

func init() {
	goloader.Register(func(loader goloader.ILoader) (err error) {
		var service ILoadService = new(LoadService)
		err = loader.Register("load.ILoadService", service)
		return
	})
}

type ILoadService interface {
	GetLoad() *Load
}
type LoadService struct {
	dao ILoadDao `load:"load.ILoadDao"`
}

func (l *LoadService) GetLoad() *Load {
	return l.dao.GetLoad()
}
