package usecase

import (
	"context"
	"crypto/rand"
	"log/slog"
	"math/big"
	"trekkstay/modules/user/constant"
	"trekkstay/pkgs/log"
	"trekkstay/utils"
)

type ForgotPasswordUseCase interface {
	ExecuteForgotPassword(ctx context.Context, email string) error
}

type forgotPasswordUseCase struct {
	mailer     Mailer
	hashAlgo   HashAlgo
	readerRepo userReaderRepository
	writerRepo userWriterRepository
}

var _ ForgotPasswordUseCase = (*forgotPasswordUseCase)(nil)

func NewForgotPasswordUseCase(mailer Mailer, hashAlgo HashAlgo,
	readerRepo userReaderRepository, writerRepo userWriterRepository) ForgotPasswordUseCase {
	return &forgotPasswordUseCase{
		mailer:     mailer,
		hashAlgo:   hashAlgo,
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func generateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	printableChars := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()_+")

	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(printableChars))))
		if err != nil {
			return "", err
		}
		password[i] = printableChars[randomIndex.Int64()]
	}

	return string(password), nil
}

func (f forgotPasswordUseCase) ExecuteForgotPassword(ctx context.Context, email string) error {
	user, err := f.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"email": email,
	})
	if err != nil {
		log.JsonLogger.Error("forgotPasswordUseCase.Execute.find_user_by_email",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrUserNotFound(err)
	}

	// Random password
	newPwd, _ := generateRandomPassword(8)

	// Hash new password
	hashedPassword, err := f.hashAlgo.HashAndSalt([]byte(newPwd))
	if err != nil {
		log.JsonLogger.Error("ExecChangePassword.hash_password",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorHashPassword(err)
	}
	user.Password = hashedPassword

	// Send email
	go func() {
		err = f.mailer.SendMail(user.Email, "Forgot password", utils.GetWorkingDirectory()+
			"/templates/forgot_password_template.html", map[string]interface{}{
			"password": newPwd,
		})

		if err != nil {
			log.JsonLogger.Error("forgotPasswordUseCase.Execute.send_mail",
				slog.Any("error", err.Error()),
				slog.String("request_id", ctx.Value("X-Request-ID").(string)),
			)
		}
	}()

	if err := f.writerRepo.UpdateUser(ctx, *user); err != nil {
		log.JsonLogger.Error("forgotPasswordUseCase.Execute.update_user",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}

	return nil
}
