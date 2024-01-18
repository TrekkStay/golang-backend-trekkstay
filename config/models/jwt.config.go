package models

type JWTConfig struct {
	JWTSecretKey       string `mapstructure:"JWT_SECRET_KEY"`
	VerifyTokenExpiry  uint   `mapstructure:"VERIFY_TOKEN_EXPIRY"`
	AccessTokenExpiry  uint   `mapstructure:"ACCESS_TOKEN_EXPIRY"`
	RefreshTokenExpiry uint   `mapstructure:"REFRESH_TOKEN_EXPIRY"`
}
