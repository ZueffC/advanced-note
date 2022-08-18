package controllers

import (
	database "todo_adv_app/database"
	models "todo_adv_app/database/models"

	"github.com/kataras/iris/v12"
)

type LinkBody struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func AddLinkController(ctx iris.Context) {
	var jsonData LinkBody

	err := ctx.ReadJSON(&jsonData)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	database.DB.Create(&models.Links{Url: jsonData.Url, Name: jsonData.Name})
	ctx.JSON(iris.Map{"status": iris.StatusOK})
}
