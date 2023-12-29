package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
}

func initLog() {
	projectSource, _ := os.Getwd()

	encoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{})
	errorLog, _ := os.Create(projectSource + "/log/error.log")
	errorCore := zapcore.NewCore(encoder, zapcore.AddSync(errorLog), zapcore.ErrorLevel)

	debugLog, _ := os.Create(projectSource + "/log/debug.log")
	debugCore := zapcore.NewCore(encoder, zapcore.AddSync(debugLog), zapcore.DebugLevel)

	tee := zapcore.NewTee(errorCore, debugCore)
	serverInstance.Logger = zap.New(tee, zap.AddCaller())
}

//func registerMvc() {
//	dir, _ := os.Getwd()
//	mvcPagePath := dir + "/src/mvc/"
//
//	set := token.NewFileSet()
//	pkgs, err := parser.ParseDir(set, mvcPagePath, nil, parser.ParseComments)
//	if err != nil {
//		serverInstance.Logger.Error("Init The System ParseDir fail~" + err.Error())
//		return
//	}
//
//	var methods []*ast.FuncDecl
//
//	for _, pkg := range pkgs {
//		// 获取包中的所以文件
//		for _, file := range pkg.Files {
//			for _, decl := range file.Decls {
//				// 检查是否是类型声明
//				if funcDecl, ok := decl.(*ast.FuncDecl); ok {
//					// 输出函数的名字
//					methods = append(methods, funcDecl)
//				}
//				//if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
//				//	for _, spec := range genDecl.Specs {
//				//		if typeDesc, ok := spec.(*ast.TypeSpec); ok {
//				//			if _, isStruct := typeDesc.Type.(*ast.FuncType); isStruct {
//				//				structList = append(structList, typeDesc)
//				//			}
//				//		}
//				//	}
//				//}
//			}
//		}
//	}
//
//	if len(methods) == 0 {
//		serverInstance.Logger.Warn("The controller in package of the size is zero~")
//		return
//	}
//
//	for _, spec := range methods {
//		methodName := spec.Name.Name
//		requestType := "GET"
//		if strings.HasPrefix(methodName, "Post") {
//			requestType = "POST"
//		}
//		handlerName := "/" + methodName
//		factory.GetServerFactoryInstance().App.Handle(requestType, handlerName, func(context *context.Context) {
//
//		})
//		//if funcType, ok := spec.Type.(*ast.FuncType); ok {
//		//for _, field := range structType.Fields.List {
//		//	for _, name := range field.Names {
//		//		fmt.Println("method name :  ", name)
//		//	}
//		//}
//		//	if funcDecl, ok := decl.(*ast.FuncDecl); ok {
//		//		// 这里可以处理函数信息，例如 funcDecl.Name
//		//		fmt.Printf("Function Name: %s\n", funcDecl.Name.Name)
//		//	}
//		//}
//		//structType := reflect.TypeOf(i)
//		//for j := 0; j < structType.NumMethod(); j++ {
//		//	method := structType.Method(j)
//		//	handlerName := "/" + method.Name
//		//	factory.GetServerFactoryInstance().App.Handle("GET", handlerName, func(context *context.Context) {
//		//		method.Func.Call([]reflect.Value{reflect.ValueOf(structType), reflect.ValueOf(context)})
//		//	})
//		//}
//	}
//}
