package user

import (
	"github.com/kataras/iris/v12/context"
	"go.uber.org/zap"
	"iris-study/src/config"
	"iris-study/src/module/factory"
	"iris-study/src/service"
)

var logger *zap.Logger
var userController = &UserController{userService: service.GetUserService()}

func init() {
	logger = config.GetLog()

	instance := factory.GetServerFactoryInstance()

	instance.App.Handle("GET", "/user/getUserName", func(context *context.Context) {
		userController.getUserName(context)
	})
	instance.App.Handle("POST", "/user/saveUser", func(context *context.Context) {
		userController.saveUser(context)
	})
}
