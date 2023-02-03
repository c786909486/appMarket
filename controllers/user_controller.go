package controllers

import (
	"appMarket/datasource"
	"appMarket/models"
	"appMarket/myUtils"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Ctx iris.Context
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	// m.Ctx.Header("Access-Control-Allow-Origin", "*")
	// if m.Ctx.Method() == "OPTIONS" {

	// 	m.Ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")

	// 	m.Ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")

	// 	m.Ctx.StatusCode(204)

	// 	return

	// }

	// m.Ctx.Next()
}

// @Summary 获取用户信息
func (uc *UserController) GetInfo() mvc.Result {
	data := uc.Ctx.URLParam("code")
	fmt.Println(data)
	return mvc.Response{
		Object: map[string]interface{}{
			"errorCode": 1,
			"data":      "12311",
		},
	}
}

func (cc *UserController) PostLogin() mvc.Result {
	x := datasource.InstanceMaster()

	var username = cc.Ctx.FormValue("username")
	var password = cc.Ctx.FormValue("password")
	if username == "" {
		return mvc.Response{
			Object: &ResponseReturn{Code: 1, Message: "用户名不可为空"},
		}
	}

	if password == "" {
		return mvc.Response{
			Object: &ResponseReturn{Code: 1, Message: "密码不可为空"},
		}
	}

	// sql := "SELECT * FROM user_model WHERE user_account = \"" + username + "\";"
	user := &models.UserModel{UserAccount: username}
	data, _ := x.Get(user)
	if !data {
		return mvc.Response{
			Object: &ResponseReturn{Code: 1, Message: "账号错误"},
		}
	}

	md5Pass := myUtils.Md5V2(password)
	if md5Pass != user.UserPassword {
		return mvc.Response{
			Object: &ResponseReturn{Code: 1, Message: "密码错误"},
		}
	}

	return mvc.Response{
		Object: &ResponseReturn{Message: "登录成功", Data: user},
	}

}

type ResponseReturn struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	List    interface{} `json:"list"`
}
