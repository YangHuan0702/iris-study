package config

import "github.com/kataras/iris/v12"

type HttpResp struct {
	Code int
	Msg  string
	Data interface{}
	Obj  interface{}
}

func Error(code int, msg string) HttpResp {
	return HttpResp{Code: code, Msg: msg}
}

func Panic(code int) HttpResp {
	return HttpResp{Code: code}
}

func Success(data interface{}) HttpResp {
	return HttpResp{Data: data, Code: iris.StatusOK}
}

func Finish() HttpResp {
	return HttpResp{Code: iris.StatusOK}
}
