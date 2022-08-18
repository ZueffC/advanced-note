package main

import (
	controllers "todo_adv_app/controllers"
	database "todo_adv_app/database"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	database.ConnectAndCreatre()

	app.Use(iris.Compression)
	app.RegisterView(iris.Handlebars("./views", ".html").Reload(true))
	app.HandleDir("/static", iris.Dir("./static"))

	app.Get("/", controllers.IndexController)
	app.Post("/add-link", controllers.AddLinkController)

	app.Listen(":49494")
}
