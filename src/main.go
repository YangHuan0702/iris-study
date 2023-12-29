package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"iris-study/src/config"
	"iris-study/src/module/factory"
)

import _ "iris-study/src/mvc/user"

func main() {
	fmt.Println("========================START Iris-Study==========================")
	fmt.Println("========================Current Mvc Finish==========================")
	instance := factory.GetServerFactoryInstance()
	logger := config.GetLog()

	err := instance.App.Run(iris.Addr(":8080"))
	if err != nil {
		logger.Error("Application Start Error: " + err.Error())
	}
}
