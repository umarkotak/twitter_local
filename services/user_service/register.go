package user_service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/contracts/request_contract"
	"github.com/umarkotak/twitter_local/models"
	"github.com/umarkotak/twitter_local/repositories/user_repository"
)

func Register(ctx context.Context, params request_contract.UserRegister) error {
	// Validate params
	err := validator.New().Struct(params)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	// Process
	_, err = user_repository.Create(ctx, models.User{
		Name:     params.Name,
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
		PhotoUrl: params.PhotoUrl,
		About:    sql.NullString{params.About, true},
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	// TODO: Kirim notifikasi selamat datang ke email ketika user sudah mendaftar

	return nil
}
