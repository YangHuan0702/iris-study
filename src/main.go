package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"iris-study/src/config"
	"iris-study/src/module/factory"
)

import _ "iris-study/src/mvc/user"
import _ "iris-study/src/interceptor"

func main() {
	fmt.Println("========================START Iris-Study==========================")
	fmt.Println("========================Current Mvc Finish==========================")
	instance := factory.GetServerFactoryInstance()
	logger := config.GetLog()

	err := instance.App.Run(iris.Addr(":8080"))

	instance.App.OnErrorCode(iris.StatusNotFound, func(context *context.Context) {
		context.JSON(config.Panic(iris.StatusNotFound))
	})

	if err != nil {
		logger.Error("Application Start Error: " + err.Error())
	}
}
