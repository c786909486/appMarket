package main

import (
	"appMarket/config"
	"appMarket/controllers"
	"appMarket/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"github.com/iris-contrib/swagger/v12"              // swagger middleware for Iris
	"github.com/iris-contrib/swagger/v12/swaggerFiles" // swagger embed files
)

// swagger middleware for Iris
// swagger embed files

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /vue
func main() {

	app := iris.New()

	// app.HandleDir("/web", iris.Dir("./web"), iris.DirOptions{IndexName: "/index.html", Compress: true})

	mvc.Configure(app.Party("/user"), userMvc)

	con := config.InitAppConfig()
	initModel()

	config := swagger.Config{
		// The url pointing to API definition.
		URL:         "http://localhost:" + con.Port + "/swagger/doc.json",
		DeepLinking: true,
	}
	swaggerUI := swagger.CustomWrapHandler(&config, swaggerFiles.Handler)

	// Register on http://localhost:8080/swagger
	app.Get("/swagger", swaggerUI)
	// And the wildcard one for index.html, *.js, *.css and e.t.c.
	app.Get("/swagger/{any:path}", swaggerUI)

	app.Run(iris.Addr(":" + con.Port))
}

func userMvc(app *mvc.Application) {
	app.Handle(new(controllers.UserController))
}

func initModel() {
	models.InitUser()
	models.InitChannel()
	models.InitApp()
}
