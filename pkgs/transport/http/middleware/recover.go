package middleware

import (
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
			if err := recover(); err != nil {
				//logger.ErrorF("Panic occurred: %+v\n", err)
				log.JsonLogger.Error("Panic occurred: %+v\n", err)

				c.Header("Content-Type", "application/json")

				if appErr, ok := err.(*res.ErrorResponse); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err)
				}

				var appErr = ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
