package usecase

import (
	"context"
	"trekkstay/modules/user/constant"
	"trekkstay/modules/user/domain/entity"
	"trekkstay/utils"
)

type CreateUserUseCase interface {
	ExecCreateUser(ctx context.Context, userEntity entity.UserEntity) error
}

type createUserUseCaseImpl struct {
	readerRepo userReaderRepository
	writerRepo userWriterRepository
}

var _ CreateUserUseCase = (*createUserUseCaseImpl)(nil)

func NewCreateUserUseCase(readerRepo userReaderRepository, writerRepo userWriterRepository) CreateUserUseCase {
	return &createUserUseCaseImpl{
		readerRepo: readerRepo,
		writerRepo: writerRepo,
	}
}

func (useCase createUserUseCaseImpl) ExecCreateUser(ctx context.Context, userEntity entity.UserEntity) error {
	// Check if user already exists
	user, err := useCase.readerRepo.FindUserByCondition(ctx, map[string]interface{}{
		"email": userEntity.Email,
	})
	if user != nil {
		return constant.ErrorEmailAlreadyExists(err)
	}

	// Hash password
	hashedPassword, err := utils.HashAndSalt([]byte(userEntity.Password))
	if err != nil {
		return constant.ErrorInternalServerError(err)
	}
	userEntity.Password = hashedPassword

	// Insert user
	err = useCase.writerRepo.InsertUser(ctx, userEntity)
	if err != nil {
		return constant.ErrorInternalServerError(err)
	}

	return nil
}
