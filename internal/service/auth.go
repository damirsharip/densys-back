package service

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/khanfromasia/densys/admin/internal/entity"
	"github.com/khanfromasia/densys/admin/internal/pkg/jwt"
	"github.com/khanfromasia/densys/admin/internal/storage/pgstorage"
	"github.com/pkg/errors"
)

// SignIn performs the sign in operation Service.
func (s *Service) SignIn(ctx context.Context, email string, password string) (entity.User, entity.Token, error) {
	var user entity.User

	maker, errJ := jwt.NewJWTMaker(s.cfg.Token.SymmetricKey)

	if errJ != nil {
		return entity.User{}, entity.Token{}, errJ
	}

	err := s.storage.ExecTX(ctx, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	}, func(queries *pgstorage.Queries) error {
		var err error

		user, err = queries.UserGetByEmail(ctx, email)

		if err != nil {
			return errors.Wrap(err, "[Service.SignIn] failed to get user by email")
		}

		return nil
	})

	if err != nil {
		return entity.User{}, entity.Token{}, errors.Wrap(err, "[Service.SignIn] failed to sign in")
	}

	//if user.Role != entity.RoleAdmin {
	//	return entity.User{}, entity.Token{}, errors.New("[Service.SignIn] user is not admin")
	//}

	if err = comparePassword(user.Password, password); err != nil {
		return entity.User{}, entity.Token{}, errors.Wrap(err, "[Service.SignIn] password does not match")
	}

	accessToken, _, errJ := maker.CreateToken(user.Email, entity.RoleAdmin, user.ID, 24*time.Hour)

	if errJ != nil {
		return entity.User{}, entity.Token{}, errors.Wrap(errJ, "[Service.SignIn] failed to create access token")
	}

	refreshToken, _, errJ := maker.CreateToken(user.Email, entity.RoleAdmin, user.ID, 24*7*time.Hour)

	if errJ != nil {
		return entity.User{}, entity.Token{}, errors.Wrap(errJ, "[Service.SignIn] failed to create refresh token")
	}

	token := entity.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return user, token, nil
}

// TokenVerify performs the token verification operation Service.
func (s *Service) TokenVerify(_ context.Context, token string) (jwt.Payload, error) {
	maker, errJ := jwt.NewJWTMaker(s.cfg.Token.SymmetricKey)

	if errJ != nil {
		return jwt.Payload{}, errJ
	}

	payload, errJ := maker.VerifyToken(token)

	if errJ != nil {
		return jwt.Payload{}, errJ
	}

	if payload == nil {
		return jwt.Payload{}, errors.New("[Service.TokenVerify] payload is nil")
	}

	return *payload, nil
}
