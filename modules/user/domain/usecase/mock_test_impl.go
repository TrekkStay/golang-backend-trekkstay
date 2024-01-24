package usecase

import (
	"context"
	"errors"
	"trekkstay/core"
	"trekkstay/modules/user/domain/entity"
)

// mockTokenProvider implements TokenProvider interface
type mockTokenProvider struct{}

func (m mockTokenProvider) Generate(payload map[string]interface{}, expiry int) (map[string]interface{}, error) {
	return map[string]interface{}{
		"token": "token_generated",
	}, nil
}

// mockHashAlgo implements HashAlgo interface
type mockHashAlgo struct {
}

func (m mockHashAlgo) HashAndSalt(pwd []byte) (string, error) {
	return "1234567890", nil
}

func (m mockHashAlgo) ComparePasswords(hashedPwd string, plainPwd []byte) error {
	return nil
}

// mockUserWriterRepository implements userWriterRepository interface
type mockUserWriterRepository struct{}

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

	if userEntity.ID == "1234567890" {
		return nil
	}

	return errors.New("user not found")
}

// mockUserReaderRepository implements userReaderRepository interface
type mockUserReaderRepository struct {
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

	if condition["id"] == "1234567890" {
		return &entity.UserEntity{
			Entity: core.Entity{
				ID: "1234567890",
			},
			Password: "password",
		}, nil
	}

	return nil, errors.New("user not found")
}

// mockMailer implements Mailer interface
type mockMailer struct{}

func (m mockMailer) SendMail(to, subject, templatePath string, data interface{}) error {
	panic("implement me")
}
