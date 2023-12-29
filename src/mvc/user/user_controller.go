package user

import (
	"github.com/kataras/iris/v12"
	"iris-study/src/domain"
	"iris-study/src/service"
)

type UserController struct {
	userService *service.UserService
}

func (userController *UserController) getUserName(ctx iris.Context) {
	_, err := ctx.HTML("<h1> Hello from /contact </h1>")
	if err != nil {
		logger.Error("userController GetUserName error:" + err.Error())
	}
}

// SaveUser 保存User信息
func (userController *UserController) saveUser(ctx iris.Context) {
	user := &domain.User{}
	if err := ctx.ReadJSON(user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Failed to read JSON"})
		return
	}
	userController.userService.Insert(user)
}
