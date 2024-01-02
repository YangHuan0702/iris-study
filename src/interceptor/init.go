package interceptor

import "iris-study/src/module/factory"

func init() {
	app := factory.GetServerFactoryInstance().App
	app.Use(ruleInterceptor)
}
