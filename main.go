package main

import (
	"github.com/Niclausse/iris-swagger-demo/handles"
	"github.com/kataras/iris/v12"
	"net/http"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"

	_ "github.com/Niclausse/iris-swagger-demo/docs" // docs is generated by Swag CLI, you have to import it.
)

// @title Iris swagger 使用示例
// @version 1.0
// @description 这是Iris Swaggo 的使用示例.

// @contact.name 林鹏
// @contact.url http://www.swagger.io/support
// @contact.email iterlp@163.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	app := iris.New()

	config := &swagger.Config{
		URL: "http://localhost:8080/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	v1 := app.Party("/api/v1")
	{
		v1.Handle(http.MethodGet, "/hello", handles.SayHello)
		v1.Handle(http.MethodGet, "/getUser", handles.PostInfo)
	}

	app.Run(iris.Addr(":8080"))
}