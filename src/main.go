package main

import "github.com/kataras/iris/v12"


func main() {
	app := iris.New()
	app.Run(iris.Addr(":8080"));
}