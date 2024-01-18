package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	res "trekkstay/core/response"
	"trekkstay/pkgs/log"
)

func ErrInternal(err error) *res.ErrorResponse {
	return res.NewErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with the server",
		"ERR_INTERNAL",
	)
}

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

				appErr = ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
