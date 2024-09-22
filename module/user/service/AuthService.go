package service

import (
	"context"
	"errors"
	contract "go-clean-architecture/module/user/contract/repository"
	"go-clean-architecture/module/user/model"
	"go-clean-architecture/module/user/request"
	"go-clean-architecture/module/user/response"
	"go-clean-architecture/utils"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepository contract.UserRepository
}

// public functions

func NewAuthService(userRepository contract.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Signup(c context.Context, req *request.SignupRequest) (response.SignupResponse, error) {
	_, err := s.GetUserByEmail(c, req.Email)

	if err != nil {
		return response.SignupResponse{}, errors.New("user already exists with the given email")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return response.SignupResponse{}, err
	}

	req.Password = string(encryptedPassword)

	user := model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	err = s.Create(c, &user)
	if err != nil {
		return response.SignupResponse{}, err
	}

	accessTokenExpiry, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	refreshTokenExpiry, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))

	accessToken, err := s.CreateAccessToken(&user, os.Getenv("ACCESS_TOKEN_SECRET"), accessTokenExpiry)
	if err != nil {
		return response.SignupResponse{}, err
	}

	refreshToken, err := s.CreateRefreshToken(&user, os.Getenv("ACCESS_TOKEN_SECRET"), refreshTokenExpiry)
	if err != nil {
		return response.SignupResponse{}, err
	}

	return response.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Login(c context.Context, req *request.LoginRequest) (response.LoginResponse, error) {
	user, err := s.GetUserByUsernameOrEmail(c, req.UsernameOrEmail)

	if err != nil {
		return response.LoginResponse{}, errors.New("Your credentials was wrong")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return response.LoginResponse{}, errors.New("Your credentials was wrong")
	}

	accessTokenExpiry, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY_HOUR"))
	refreshTokenExpiry, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRY_HOUR"))

	accessToken, err := s.CreateAccessToken(&user, os.Getenv("ACCESS_TOKEN_SECRET"), accessTokenExpiry)
	if err != nil {
		return response.LoginResponse{}, err
	}

	refreshToken, err := s.CreateRefreshToken(&user, os.Getenv("ACCESS_TOKEN_SECRET"), refreshTokenExpiry)
	if err != nil {
		return response.LoginResponse{}, err
	}

	return response.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) Create(c context.Context, user *model.User) error {
	return s.userRepository.Create(c, user)
}

func (s *AuthService) GetUserByUsernameOrEmail(c context.Context, usernameOrEmail string) (model.User, error) {
	return s.userRepository.GetByUsernameOrEmail(c, usernameOrEmail)
}

func (s *AuthService) GetUserByEmail(c context.Context, email string) (model.User, error) {
	return s.userRepository.GetByEmail(c, email)
}

func (s *AuthService) GetUserByID(c context.Context, userID string) (model.Profile, error) {
	user, err := s.userRepository.GetByID(c, userID)

	if err != nil {
		return model.Profile{}, err
	}

	return model.Profile{Username: user.Username, Email: user.Email}, nil
}

func (s *AuthService) CreateAccessToken(user *model.User, secret string, expiry int) (string, error) {
	return utils.CreateAccessToken(user, secret, expiry)
}

func (s *AuthService) CreateRefreshToken(user *model.User, secret string, expiry int) (string, error) {
	return utils.CreateRefereshToken(user, secret, expiry)
}
