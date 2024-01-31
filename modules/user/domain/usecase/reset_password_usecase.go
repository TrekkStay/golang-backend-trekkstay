package usecase

import (
	"context"
	"log/slog"
	"trekkstay/modules/user/constant"
	"trekkstay/pkgs/log"
	"trekkstay/utils"
)

type ResetPasswordUseCase interface {
	ExecuteResetPassword(ctx context.Context, email string) error
}

type resetPasswordUseCase struct {
	mailer     Mailer
	hashAlgo   HashAlgo
	readerRepo userReaderRepository
	writerRepo userWriterRepository
}

var _ ResetPasswordUseCase = (*resetPasswordUseCase)(nil)

func NewResetPasswordUseCase(mailer Mailer, hashAlgo HashAlgo,
	readerRepo userReaderRepository, writerRepo userWriterRepository) ResetPasswordUseCase {
	return &resetPasswordUseCase{
		mailer:     mailer,
		hashAlgo:   hashAlgo,
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase resetPasswordUseCase) ExecuteResetPassword(ctx context.Context, email string) error {
	user, err := useCase.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		log.JsonLogger.Error("ExecuteResetPassword.find_user_by_email",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrUserNotFound(err)
	}

	// Random password
	newPwd, _ := utils.GenerateRandomPassword(8)

	// Hash new password
	hashedPassword, err := useCase.hashAlgo.HashAndSalt([]byte(newPwd))
	if err != nil {
		log.JsonLogger.Error("ExecuteResetPassword.hash_password",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorHashPassword(err)
	}
	user.Password = hashedPassword

	// Send email
	go func() {
		err = useCase.mailer.SendMail(user.Email, "Forgot password", utils.GetWorkingDirectory()+
			"/templates/forgot_password_template.html", map[string]interface{}{
			"password": newPwd,
		})

		if err != nil {
			log.JsonLogger.Error("ExecuteResetPassword.send_mail",
				slog.Any("error", err.Error()),
				slog.String("request_id", ctx.Value("X-Request-ID").(string)),
			)
		}
	}()

	if err := useCase.writerRepo.UpdateUser(ctx, *user); err != nil {
		log.JsonLogger.Error("ExecuteResetPassword.update_user",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}

	return nil
}
