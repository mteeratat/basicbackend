package auth

import (
	"net/http"
	"os"
	"phase2/customError"
	"phase2/customLog"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte("your-secret-key")

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger := customLog.NewCustomLogger(customLog.LevelInfo, os.Stdout)
		authHeader := c.Request().Header.Get("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			logger.Error("Unauthorized")
			return c.JSON(http.StatusUnauthorized, customError.NewMyError(http.StatusUnauthorized, "Unauthorized"))
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, customError.NewMyError(http.StatusUnauthorized, "Unauthorized")
			}
			return secretKey, nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, customError.NewMyError(http.StatusUnauthorized, "Unauthorized"))
		}

		return next(c)
	}
}

func GenToken(userID int) (any, int64, error) {
	expTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    expTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return nil, 0, customError.NewMyError(http.StatusUnauthorized, "JWT token signed failed")
	}

	return signedToken, expTime, nil
}
