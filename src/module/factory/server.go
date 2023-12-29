package factory

import (
	"github.com/kataras/iris/v12"
	"sync"
)

type ServerFactory struct {
	App *iris.Application
}

var mutex = sync.Mutex{}
var serverFactory *ServerFactory = nil

func GetServerFactoryInstance() *ServerFactory {
	mutex.Lock()
	if serverFactory == nil {
		serverFactory = &ServerFactory{}
		serverFactory.App = iris.New()
	}
	mutex.Unlock()
	return serverFactory
}
