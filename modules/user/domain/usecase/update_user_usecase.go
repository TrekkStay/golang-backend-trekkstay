package usecase

import (
	"context"
	"log/slog"
	"trekkstay/api/middlewares/constant"
	"trekkstay/core"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/pkgs/log"
)

type UpdateUserUseCase interface {
	ExecUpdateUser(ctx context.Context, userEntity entity.UserEntity) error
}

type updateUserUseCaseImpl struct {
	readerRepo userReaderRepository
	writerRepo userWriterRepository
}

var _ UpdateUserUseCase = (*updateUserUseCaseImpl)(nil)

func NewUpdateUserUseCase(readerRepo userReaderRepository, writerRepo userWriterRepository) UpdateUserUseCase {
	return &updateUserUseCaseImpl{
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase updateUserUseCaseImpl) ExecUpdateUser(ctx context.Context, userEntity entity.UserEntity) error {
	userID := ctx.Value(core.CurrentRequesterKeyStruct{}).(core.Requester).GetUserID()
	userEntity.ID = userID

	// Update user
	err := useCase.writerRepo.UpdateUser(ctx, userEntity)
	if err != nil {
		log.JsonLogger.Error("ExecUpdateUser.update_user",
			slog.Any("error", err.Error()),
			slog.String("request_id", ctx.Value("X-Request-ID").(string)),
		)

		return constant.ErrInternal(err)
	}

	return nil
}
