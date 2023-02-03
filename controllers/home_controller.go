package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type HomeController struct {
	Ctx iris.Context
}

func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "/web/index.html",
		Data: iris.Map{"Title": "首页"},
	}
}

func (c *HomeController) POST() mvc.Result {

	return mvc.Response{
		Object: map[string]interface{}{
			"errorCode": 1,
			"data":      "12311",
		},
	}
}
