package controllers

import (
	database "todo_adv_app/database"
	models "todo_adv_app/database/models"

	"github.com/kataras/iris/v12"
)

func IndexController(ctx iris.Context) {
	var links []models.Links
	var notes []models.Notes

	database.DB.Find(&links)
	database.DB.Find(&notes)

	ctx.ViewLayout("layouts/main")
	ctx.View("index", iris.Map{"links": links, "notes": notes})
}
