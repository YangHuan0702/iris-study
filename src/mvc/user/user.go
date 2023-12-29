package user

import "github.com/kataras/iris/v12"

func GetUserName(ctx iris.Context) {
	_, err := ctx.HTML("<h1> Hello from /contact </h1>")
	if err != nil {

	}
}
