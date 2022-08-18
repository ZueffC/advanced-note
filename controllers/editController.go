package controllers

import (
	database "todo_adv_app/database"
	models "todo_adv_app/database/models"

	"github.com/kataras/iris/v12"
)

func EditController(ctx iris.Context) {
	type urlParams struct {
		Type string `param:"type"`
	}

	var binder urlParams

	if err := ctx.ReadParams(&binder); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if binder.Type == "link" {
		editLink(ctx)
		ctx.JSON(iris.Map{"status": iris.StatusOK})
	}
}

func editLink(ctx iris.Context) {
	type linkBody struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Url      string `json:"url"`
		IsDelete int    `json:"delete"`
	}

	var jsonData linkBody

	if err := ctx.ReadJSON(&jsonData); err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	if jsonData.IsDelete == 0 {
		database.DB.Model(&models.Links{}).Where("id", jsonData.Id).Updates(models.Links{Name: jsonData.Name, Url: jsonData.Url})
	} else {
		database.DB.Where("id = ?", jsonData.Id).Delete(&models.Links{})
	}
}
