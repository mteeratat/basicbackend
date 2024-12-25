package main

import (
	"fmt"
	"net/http"
	"os"
	"phase2/customError"
	"phase2/customLog"
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
	customLog := customLog.NewCustomLogger(customLog.LevelInfo, os.Stdout)

	// e.Use(middleware.Logger())
	e.HideBanner = true
	// e.HidePort = true

	e.GET("/", func(ctx echo.Context) error {
		msg := "Hello World!"
		// log.Info(msg)
		customLog.Info(msg)
		return ctx.String(http.StatusOK, msg)
	})
	e.POST("/create", func(ctx echo.Context) error {
		var req model.Todo
		if err := ctx.Bind(&req); err != nil {
			msg := "can't bind req"
			customLog.Error(msg)
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg, err))
		}
		if err := validator.Struct(req); err != nil {
			msg := "validate failed"
			customLog.Error(msg)
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg, err))
		}
		req.ID = count
		msg := fmt.Sprintf("Create [%d]: %v %v %v", count, req.ID, *req.Title, *req.Status)
		// log.Infof(msg)
		customLog.Info(msg)
		temp = append(temp, req)
		count++
		return ctx.JSON(http.StatusOK, req)
	})
	e.GET("/getall", func(ctx echo.Context) error {
		msg := fmt.Sprintf("Get All : %v", temp)
		// log.Infof(msg)
		customLog.Info(msg)
		return ctx.JSON(http.StatusOK, temp)
	})
	e.GET("/get/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			msg := "id too much"
			customLog.Error(msg)
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := fmt.Sprintf("Get [%s] : %v", id, temp[index])
		// log.Info(msg)
		customLog.Info(msg)
		return ctx.JSON(http.StatusOK, temp[index])
	})
	e.PUT("/update/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			msg := "id too much"
			customLog.Error(msg)
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := fmt.Sprintf("Update [%s] : %v", id, temp[index])
		// log.Infof(msg)
		customLog.Info(msg)

		var req model.Todo
		if err := ctx.Bind(&req); err != nil {
			msg := "can't bind req"
			customLog.Error(msg)
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg, err))
		}
		msg = fmt.Sprintf("Using [%d]: %v", index, req)
		// log.Info(msg)
		customLog.Info(msg)

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
			msg := "id too much"
			customLog.Error(msg)
			return ctx.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := fmt.Sprintf("Delete [%s] : %v", id, temp[index])
		// log.Infof(msg)
		customLog.Info(msg)
		temp = append(temp[:index], temp[index+1:]...)
		return ctx.JSON(http.StatusOK, temp)
	})
	log.Fatal(e.Start(":8080"))
}
