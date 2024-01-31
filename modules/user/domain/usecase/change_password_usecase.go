package usecase

import (
	"context"
	"log/slog"
	"trekkstay/core"
	"trekkstay/modules/user/constant"
	"trekkstay/pkgs/log"
)

type ChangePasswordUseCase interface {
	ExecChangePassword(ctx context.Context, oldPwd, newPwd string) error
}

type changePasswordUseCaseImpl struct {
	hashAlgo   HashAlgo
	readerRepo userReaderRepository
	writerRepo userWriterRepository
}

var _ ChangePasswordUseCase = (*changePasswordUseCaseImpl)(nil)

func NewChangePasswordUseCase(hashAlgo HashAlgo, readerRepo userReaderRepository,
	writerRepo userWriterRepository) ChangePasswordUseCase {
	return &changePasswordUseCaseImpl{
		hashAlgo:   hashAlgo,
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase changePasswordUseCaseImpl) ExecChangePassword(ctx context.Context, oldPwd, newPwd string) error {
	// Find user by id
	user, err := useCase.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"id": ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID(),
	})
	if err != nil {
		log.JsonLogger.Error("ExecChangePassword.find_user_by_id",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrUserNotFound(err)
	}

	// Compare old password
	if err := useCase.hashAlgo.ComparePasswords(user.Password, []byte(oldPwd)); err != nil {
		log.JsonLogger.Error("ExecChangePassword.compare_password",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorWrongPassword(err)
	}

	// Hash new password
	hashedPassword, err := useCase.hashAlgo.HashAndSalt([]byte(newPwd))
	if err != nil {
		log.JsonLogger.Error("ExecChangePassword.hash_password",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorHashPassword(err)
	}
	user.Password = hashedPassword

	// Update user
	if err := useCase.writerRepo.UpdateUser(ctx, *user); err != nil {
		log.JsonLogger.Error("ExecChangePassword.update_user",
			slog.String("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrorInternalServerError(err)
	}

	return nil
}
