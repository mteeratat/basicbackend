package main

import (
	"net/http"
	"phase2/customError"
	"phase2/model"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	temp := []model.Todo{}
	count := 0
	e := echo.New()
	validator := validator.New()

	// e.Use(middleware.Logger())
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", func(ctx echo.Context) error {
		log.Info("Hello World!")
		return ctx.String(http.StatusOK, "Hello World!")
	})
	e.POST("/create", func(ctx echo.Context) error {
		var req model.Todo
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "can't bind req", err))
		}
		if err := validator.Struct(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "validate failed", err))
		}
		req.ID = count
		log.Infof("Create [%d]: %v %v %v", count, req.ID, *req.Title, *req.Status)
		temp = append(temp, req)
		count++
		return ctx.JSON(http.StatusOK, req)
	})
	e.GET("/getall", func(ctx echo.Context) error {
		log.Infof("Get All : %v", temp)
		return ctx.JSON(http.StatusOK, temp)
	})
	e.GET("/get/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "id too much"))
		}
		log.Infof("Get [%s] : %v", id, temp[index])
		return ctx.JSON(http.StatusOK, temp[index])
	})
	e.PUT("/update/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "id too much"))
		}
		log.Infof("Update [%s] : %v", id, temp[index])

		var req model.Todo
		if err := ctx.Bind(&req); err != nil {
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "can't bind req", err))
		}
		log.Infof("Using [%d]: %v", index, req)

		if req.Title != nil {
			temp[index].Title = req.Title
		}
		if req.Status != nil {
			temp[index].Status = req.Status
		}

		return ctx.JSON(http.StatusOK, temp[index])
	})
	e.DELETE("delete/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "id too much"))
		}
		log.Infof("Delete [%s] : %v", id, temp[index])
		temp = append(temp[:index], temp[index+1:]...)
		return ctx.JSON(http.StatusOK, temp)
	})
	log.Fatal(e.Start(":8080"))
}
