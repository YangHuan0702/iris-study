package interceptor

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func ruleInterceptor(ctx iris.Context) {
	route := ctx.GetCurrentRoute()
	fmt.Println("Call " + route.Path() + " Before...")
	ctx.Next()
	fmt.Println("Call " + route.Path() + " After...")
}
