package config

import (
	"github.com/kataras/iris/v12/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go/ast"
	"go/parser"
	"go/token"
	"iris-study/src/module/factory"
	"os"
	"reflect"
	"sync"
)

var serverInstance *ServerInstance = nil
var mutex = sync.Mutex{}

type ServerInstance struct {
	Logger *zap.Logger
}

// 获取Logger对象
func GetLog() *zap.Logger {
	mutex.Lock()
	if serverInstance == nil {
		serverInstance = &ServerInstance{}
		initLog()
	}
	mutex.Unlock()

	return serverInstance.Logger
}

func init() {

	mutex.Lock()
	if serverInstance == nil {
		serverInstance = &ServerInstance{}
	}
	mutex.Unlock()

	// init log
	initLog()

	// configuration mvc register
	registerMvc()
}

func initLog() {
	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	errorLog, _ := os.Create("../../log/error.log")
	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorLog), zapcore.ErrorLevel)

	debugLog, _ := os.Create("../../log/debug.log")
	debugCore := zapcore.NewCore(encoder, zapcore.AddSync(debugLog), zapcore.DebugLevel)

	tee := zapcore.NewTee(errorCore, debugCore)
	serverInstance.Logger = zap.New(tee, zap.AddCaller())
}

func registerMvc() {
	mvcPagePath := "../mvc/"

	set := token.NewFileSet()
	pkgs, err := parser.ParseDir(set, mvcPagePath, nil, parser.ParseComments)
	if err != nil {
		serverInstance.Logger.Error("Init The System ParseDir fail~")
	}

	var structList []*ast.TypeSpec

	for _, pkg := range pkgs {
		// 获取包中的所以文件
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				// 检查是否是类型声明
				if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
					for _, spec := range genDecl.Specs {
						if typeDesc, ok := spec.(*ast.TypeSpec); ok {
							if _, isStruct := typeDesc.Type.(*ast.StructType); isStruct {
								structList = append(structList, typeDesc)
							}
						}
					}
				}
			}
		}
	}

	if len(structList) == 0 {
		serverInstance.Logger.Warn("The controller in package of the size is zero~")
		return
	}

	for _, i := range structList {
		structType := reflect.TypeOf(i)
		for index := 0; index < structType.NumMethod(); index++ {
			method := structType.Method(index)
			handlerName := "/" + method.Name
			factory.GetServerFactoryInstance().App.Handle("POST", handlerName, func(context *context.Context) {
				method.Func.Call([]reflect.Value{reflect.ValueOf(structType), reflect.ValueOf(context)})
			})
		}
	}
}
