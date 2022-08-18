package controllers

import (
	database "todo_adv_app/database"
	models "todo_adv_app/database/models"

	"github.com/kataras/iris/v12"
)

func AddController(ctx iris.Context) {
	type urlParams struct {
		Type string `param:"type"`
	}

	var binder urlParams

	if err := ctx.ReadParams(&binder); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if binder.Type == "link" {
		createLinkRow(ctx)
		ctx.JSON(iris.Map{"status": iris.StatusOK})
		return
	} else if binder.Type == "note" {
		createNoteRow(ctx)
		ctx.JSON(iris.Map{"status": iris.StatusOK})
	}
}

func createLinkRow(ctx iris.Context) {
	type linkBody struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	var jsonData linkBody

	if err := ctx.ReadJSON(&jsonData); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	database.DB.Create(&models.Links{Url: jsonData.Url, Name: jsonData.Name})
}
func createNoteRow(ctx iris.Context) {
	type noteBody struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}

	var jsonData noteBody

	if err := ctx.ReadJSON(&jsonData); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	database.DB.Create(&models.Notes{Title: jsonData.Title, Text: jsonData.Text})
}
