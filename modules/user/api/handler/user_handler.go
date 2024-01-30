package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"trekkstay/modules/user/domain/usecase"
)

type UserHandler interface {
	HandleCreateUser(c *gin.Context)
	HandleUpdateUser(c *gin.Context)
	HandleLoginUser(c *gin.Context)
	HandleChangePassword(c *gin.Context)
	HandleResetPassword(c *gin.Context)
	HandleRefreshToken(c *gin.Context)
}

type userHandler struct {
	requestValidator      *validator.Validate
	createUserUseCase     usecase.CreateUserUseCase
	updateUserUseCase     usecase.UpdateUserUseCase
	loginUserUseCase      usecase.LoginUserUseCase
	changePasswordUseCase usecase.ChangePasswordUseCase
	resetPasswordUseCase  usecase.ResetPasswordUseCase
	refreshTokenUseCase   usecase.RefreshTokenUseCase
}

func NewUserHandler(
	requestValidator *validator.Validate,
	createUserUseCase usecase.CreateUserUseCase,
	updateUserUseCase usecase.UpdateUserUseCase,
	loginUserUseCase usecase.LoginUserUseCase,
	changePasswordUseCase usecase.ChangePasswordUseCase,
	resetPasswordUseCase usecase.ResetPasswordUseCase,
	refreshTokenUseCase usecase.RefreshTokenUseCase,
) UserHandler {
	return &userHandler{
		requestValidator:      requestValidator,
		createUserUseCase:     createUserUseCase,
		updateUserUseCase:     updateUserUseCase,
		loginUserUseCase:      loginUserUseCase,
		changePasswordUseCase: changePasswordUseCase,
		resetPasswordUseCase:  resetPasswordUseCase,
		refreshTokenUseCase:   refreshTokenUseCase,
	}
}
