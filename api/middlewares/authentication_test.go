package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"testing"
	"trekkstay/api/middlewares/constant"
	res "trekkstay/core/response"
	"trekkstay/pkgs/log"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover().(error); err != nil {
				//logger.ErrorF("Panic occurred: %+v\n", err)
				log.JsonLogger.Error("Panic occurred: %+v\n", err)

				c.Header("Content-Type", "application/json")

				var appErr *res.ErrorResponse
				if errors.As(err, &appErr) {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				appErr = constant.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}

func TestAuthentication(t *testing.T) {
	// Visit http://localhost:8080/auth-test to test

	app := gin.Default()
	app.Use(Recover())
	app.Use(Authentication())
	app.GET("/auth-test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Authenticated",
		})
	})

	err := app.Run(":8080")
	if err != nil {
		return
	}
}
