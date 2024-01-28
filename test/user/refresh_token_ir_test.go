package user

import (
	"context"
	"os"
	"testing"
	"trekkstay/config"
	"trekkstay/config/models"
	"trekkstay/core"
	"trekkstay/modules/user/domain/usecase"
	"trekkstay/pkgs/jwt"
)

func TestIRRefreshToken(t *testing.T) {
	err := os.Setenv("CONFIG_PATH", "../../env/dev.env")
	if err != nil {
		return
	}

	jwtConfig := config.LoadConfig(&models.JWTConfig{}).(*models.JWTConfig)
	jwtToken := jwt.NewJWT(jwtConfig.JWTSecretKey)

	useCase := usecase.NewRefreshTokenUseCase(
		jwtToken, jwtConfig.AccessTokenExpiry, jwtConfig.RefreshTokenExpiry)

	ctx := context.WithValue(context.Background(), "X-Request-ID", "1234567890")
	ctx = context.WithValue(ctx, core.CurrentRequesterKeyStruct{}, core.RestRequester{
		ID: "151d3f25-7c4e-4c9a-a3b8-55356ebcfbf56",
	})

	_, _, err = useCase.ExecRefreshToken(ctx)
	if err != nil {
		t.Error(err)
	}
}
