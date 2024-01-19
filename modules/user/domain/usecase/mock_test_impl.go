package usecase

import (
	"context"
	"errors"
	"trekkstay/modules/user/domain/entity"
)

type mockUserReaderRepository struct {
}

type mockUserWriterRepository struct{}

type mockJWT struct{}

type mockHashAlgo struct {
}

func (m mockHashAlgo) HashAndSalt(pwd []byte) (string, error) {
	return "1234567890", nil
}

func (m mockHashAlgo) ComparePasswords(hashedPwd string, plainPwd []byte) error {
	return nil
}

func (m mockJWT) Generate(payload map[string]interface{}, expiry int) (map[string]interface{}, error) {
	return map[string]interface{}{
		"token": "token_generated",
	}, nil
}

func (m mockUserWriterRepository) InsertUser(ctx context.Context, userEntity entity.UserEntity) error {
	if userEntity.Email == "existedemail@example.com" {
		return errors.New("email already existed")
	}

	return nil
}

func (m mockUserWriterRepository) DeleteUser(ctx context.Context, userId string) error {
	if userId == "1234567890" {
		return nil
	}

	return errors.New("user not found")
}

func (m mockUserWriterRepository) UpdateUser(ctx context.Context, userEntity entity.UserEntity) error {
	if userEntity.Email == "existedemail@example.com" {
		return nil
	}

	return errors.New("user not found")
}

func (m mockUserReaderRepository) FindUserByCondition(ctx context.Context,
	condition map[string]interface{}) (*entity.UserEntity, error) {
	if condition["email"] == "existedemail@example.com" {
		return &entity.UserEntity{
			Email:    "existedemail@example.com",
			Password: "password",
		}, nil
	}

	if condition["phone"] == "12345678900" {
		return &entity.UserEntity{
			Phone: "12345678900",
		}, nil
	}

	return nil, errors.New("user not found")
}
