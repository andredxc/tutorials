package main

import (
	"github.com/labstack/echo/v4"
	"go-api/list_repository"
	"go-api/router"
)

func main() {

	e := echo.New()
	r := router.ListRouter{ListRepository: list_repository.Repository{}}

	e.POST("/list/create", r.CreateList)
	e.POST("/list/addItem", r.AddItem)
	e.POST("/list/changeItem", r.ChangeItem)
	e.DELETE("/list/removeItem", r.RemoveItem)
	e.GET("/list/get", r.GetList)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
