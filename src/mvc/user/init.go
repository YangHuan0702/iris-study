package user

import (
	"github.com/kataras/iris/v12/context"
	"iris-study/src/module/factory"
)

func init() {
	instance := factory.GetServerFactoryInstance()

	instance.App.Handle("GET", "/user/GetUserName", func(context *context.Context) {
		GetUserName(context)
	})

}
