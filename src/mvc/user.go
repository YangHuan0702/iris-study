package mvc

import "github.com/kataras/iris/v12"

type UserMvc struct {
	userName string
	age      int
}

func (user *UserMvc) getUserName(ctx iris.Context) {
	_, err := ctx.HTML("<h1> Hello from /contact </h1>")
	if err != nil {

	}
}
