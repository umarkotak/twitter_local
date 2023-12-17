package jwt_auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/twitter_local/models"
)

var signKey = []byte("some-random-secret-singing-key")

type UserClaims struct {
	jwt.RegisteredClaims
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func GenUserAuthToken(ctx context.Context, user models.User) (string, error) {
	jwtID, _ := uuid.NewRandom()

	claims := &UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "twitter_local",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			ID:        jwtID.String(),
		},
		UserID:   user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := token.SignedString(signKey)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return "", err
	}

	return jwtToken, nil
}

func DecodeUserAuthToken(ctx context.Context, jwtToken string) (models.User, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return signKey, nil
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return models.User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = fmt.Errorf("invalid map claims")
		logrus.WithContext(ctx).Error(err)
		return models.User{}, err
	}

	return models.User{
		ID:       claims["user_id"].(string),
		Name:     claims["name"].(string),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
	}, nil
}
