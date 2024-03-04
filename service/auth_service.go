package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/luquxSentinel/parkingspace/api/token"
	"github.com/luquxSentinel/parkingspace/storage"
	"github.com/luquxSentinel/parkingspace/types"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationService interface {
	SignUp(ctx context.Context, data *types.SignUpData) error
	SignIn(ctx context.Context, data *types.SignInData) (*types.User, *string, error)
	ChangePassword(ctx context.Context, newPassword string) error
}

type authenticationService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) *authenticationService {
	return &authenticationService{
		storage: storage,
	}
}

func (s *authenticationService) SignUp(ctx context.Context, data *types.SignUpData) error {
	// check if email exists
	count, err := s.storage.CountUserEmail(ctx, data.Email)
	if err != nil {
		return errors.New("failed to verify email")
	}

	if count > 0 {
		return errors.New("email already in use")
	}

	// check of phone number exists
	count, err = s.storage.CountUserPhoneNumber(ctx, data.PhoneNumber)
	if err != nil {
		return errors.New("fail to verify phone number")
	}

	if count > 0 {
		return errors.New("phone number already in use")
	}

	// hash password
	pwd, err := hashPassword(data.Password)
	if err != nil {
		return errors.New("failed to create user")
	}
	// create user from data
	user := &types.User{
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Email:          data.Email,
		PhoneNumber:    data.PhoneNumber,
		IdentityNumber: data.IdentityNumber,
		CreatedAt:      time.Now().UTC(),
		Password:       pwd,
	}

	// insert user into store
	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *authenticationService) SignIn(ctx context.Context, data *types.SignInData) (*types.User, *string, error) {
	//  get user by email from store
	user, err := s.storage.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, errors.New("wrong email or password")
	}

	// validate password
	err = verifyPassword(user.Password, data.Password)
	if err != nil {
		return nil, nil, errors.New("wrong email or password")
	}

	//  generate jwt
	token, err := token.NewJwt(user.UID, user.Email)
	if err != nil {
		// we are definitely fucked
		log.Panicf("error occured while generating jwt token : %v", err)
		return nil, nil, errors.New("unexpacted error occured. please try again")

	}

	// return user & jwt
	return user, token, nil
}

func ChangePassword(ctx context.Context, newPassword string) error {
	// TODO implement change password func

	return nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(b), err
}

func verifyPassword(hash, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
