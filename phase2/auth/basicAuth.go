package auth

import (
	"encoding/base64"
	"net/http"
	"os"
	"phase2/customError"
	"phase2/customLog"
	"strings"

	"github.com/labstack/echo/v4"
)

func BasicAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := customLog.NewCustomLogger(customLog.LevelInfo, os.Stdout)
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Basic ") {
			logger.Error("Unauthorized")
			return c.JSON(http.StatusUnauthorized, customError.NewMyError(http.StatusUnauthorized, "Unauthorized"))
		}

		encodedCredentials := strings.TrimPrefix(authHeader, "Basic ")

		decodedCredentials, err := base64.StdEncoding.DecodeString(encodedCredentials)
		if err != nil {
			logger.Error("Invalid Authorization header")
			return c.JSON(http.StatusUnauthorized, customError.NewMyError(http.StatusUnauthorized, "Invalid Authorization header"))
		}

		credentials := strings.SplitN(string(decodedCredentials), ":", 2)
		if len(credentials) != 2 || !isValidUser(credentials[0], credentials[1]) {
			logger.Error("Unauthorized")
			return c.JSON(http.StatusUnauthorized, customError.NewMyError(http.StatusUnauthorized, "Unauthorized"))
		}

		logger.Info(authHeader)

		return next(c)
	}
}

func isValidUser(username, password string) bool {
	// Basic YWRtaW46cGFzc3dvcmQ=
	return username == "admin" && password == "password"
}
