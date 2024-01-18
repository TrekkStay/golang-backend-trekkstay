package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
	"trekkstay/api/middlewares/constant"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/pkgs/jwt"
)

// extractTokenFromHeader extracts a token from the given header string.
//
// It takes a header string as a parameter and returns a pointer to a string and an error.
func extractTokenFromHeader(header string) (*string, error) {
	parts := strings.Fields(header)

	if len(parts) == 0 {
		return nil, constant.ErrMissingToken(nil)
	}

	// Check if the header is in the correct format Bearer {token}
	if len(parts) < 2 || parts[0] != "Bearer" || parts[1] == "" {
		return nil, constant.ErrInvalidToken(nil)
	}

	return &parts[1], nil
}

// Authentication returns a middleware function that performs authentication.
//
// The middleware function takes a gin.Context parameter and validates the
// token found in the "Authorization" header of the request. If the token is
// valid, it sets the requester information in the context and calls the next
// middleware function.
//
// Parameters:
// - ctx: A pointer to a gin.Context object.
//
// Returns:
// - A function that takes a *gin.Context parameter.
func Authentication() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
		jwtProvider := jwt.NewJWT(jwtConfig.JWTSecretKey)

		// Extract token from header
		token, err := extractTokenFromHeader(c.GetHeader("Authorization"))
		if err != nil {
			panic(err)
		}

		// Validate token
		payload, err := jwtProvider.Validate(*token)
		if err != nil {
			panic(constant.ErrInvalidToken(err))
		}

		// Set requester information in context
		c.Set(core.CurrentRequesterKey, core.RestRequester{
			Id: payload.UserId,
		})

		c.Next()
	}
}
