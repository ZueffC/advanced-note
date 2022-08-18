package controllers

import (
	database "todo_adv_app/database"
	models "todo_adv_app/database/models"

	"github.com/kataras/iris/v12"
)

func GetController(ctx iris.Context) {
	type urlParams struct {
		Type string `param:"type"`
		Id   int    `param:"id"`
	}

	var binder urlParams

	if err := ctx.ReadParams(&binder); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if binder.Type == "link" {
		linkData := getLinkById(binder.Id)
		ctx.JSON(iris.Map{"data": linkData})
	}
}

func getLinkById(id int) models.Links {
	var linkData models.Links
	database.DB.Find(&linkData, id)

	return linkData
}
