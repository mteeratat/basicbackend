package main

import (
	"fmt"
	"net/http"
	"os"
	"phase2/auth"
	"phase2/customError"
	"phase2/customLog"
	"phase2/model"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	temp := []model.Todo{}
	count := 0
	e := echo.New()
	validator := validator.New()
	logger := customLog.NewCustomLogger(customLog.LevelInfo, os.Stdout)

	// e.Use(middleware.Logger())
	// e.Use(auth.BasicAuthMiddleware)
	e.HideBanner = true
	// e.HidePort = true

	e.GET("/", func(c echo.Context) error {
		msg := "Hello World!"
		// log.Info(msg)
		logger.Info(msg)
		return c.String(http.StatusOK, msg)
	})
	e.POST("/create", func(c echo.Context) error {
		var req model.Todo
		if err := c.Bind(&req); err != nil {
			msg := "can't bind req"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		if err := validator.Struct(req); err != nil {
			msg := "validate failed"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		req.ID = count
		msg := fmt.Sprintf("Create [%d]: %v %v %v", count, req.ID, *req.Title, *req.Status)
		// log.Infof(msg)
		logger.Info(msg)
		temp = append(temp, req)
		count++
		return c.JSON(http.StatusOK, req)
	}, auth.JWTMiddleware)
	e.GET("/getall", func(c echo.Context) error {
		msg := fmt.Sprintf("Get All : %v", temp)
		// log.Infof(msg)
		logger.Info(msg)
		return c.JSON(http.StatusOK, temp)
	})
	e.GET("/get/:id", func(c echo.Context) error {
		id := c.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			msg := "id too much"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := fmt.Sprintf("Get [%s] : %v", id, temp[index])
		// log.Info(msg)
		logger.Info(msg)
		return c.JSON(http.StatusOK, temp[index])
	})
	e.PUT("/update/:id", func(c echo.Context) error {
		id := c.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			msg := "id too much"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := fmt.Sprintf("Update [%s] : %v", id, temp[index])
		// log.Infof(msg)
		logger.Info(msg)

		var req model.Todo
		if err := c.Bind(&req); err != nil {
			msg := "can't bind req"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg = fmt.Sprintf("Using [%d]: %v", index, req)
		// log.Info(msg)
		logger.Info(msg)

		if req.Title != nil {
			temp[index].Title = req.Title
		}
		if req.Status != nil {
			temp[index].Status = req.Status
		}

		return c.JSON(http.StatusOK, temp[index])
	}, auth.JWTMiddleware)
	e.DELETE("/delete/:id", func(c echo.Context) error {
		id := c.Param("id")
		index, _ := strconv.Atoi(id)
		if index >= len(temp) || index < 0 {
			msg := "id too much"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := fmt.Sprintf("Delete [%s] : %v", id, temp[index])
		// log.Infof(msg)
		logger.Info(msg)
		temp = append(temp[:index], temp[index+1:]...)
		return c.JSON(http.StatusOK, temp)
	}, auth.JWTMiddleware)
	e.POST("/genToken", func(c echo.Context) error {
		var req model.JWTtoken
		if err := c.Bind(&req); err != nil {
			msg := "can't bind req"
			logger.Error(msg)
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, msg))
		}
		msg := "Generating JWT token"
		logger.Info(msg)

		signedToken, expTime, err := auth.GenToken(req.UserID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, customError.NewMyError(http.StatusBadRequest, "can't gen JWT"))
		}
		msg = fmt.Sprintf("Signed JWT token: %s\nExpired at %s", signedToken, time.Unix(expTime, 0))
		logger.Info(msg)
		return c.String(http.StatusOK, msg)
	}, auth.BasicAuthMiddleware)

	log.Fatal(e.Start(":8080"))
}
