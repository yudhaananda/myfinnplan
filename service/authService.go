package service

import (
	"errors"
	"math/rand"
	"myfinnplan/entity"
	"myfinnplan/input"
	"myfinnplan/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(input input.UserInput) (entity.User, error)
	Login(input input.LoginInput) (entity.User, error)
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *authService {
	return &authService{userRepository}
}

func (s *authService) RegisterUser(input input.UserInput) (entity.User, error) {

	checkUser, err := s.userRepository.FindByUserName(input.UserName)

	if err != nil {
		return entity.User{}, errors.New("error find user")
	}

	if len(checkUser) != 0 {
		return entity.User{}, errors.New("UserName sudah pernah diinputkan")
	}

	key := rand.Intn(9)
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), key)
	if err != nil {
		return entity.User{}, errors.New("error encrypt password")
	}

	user := entity.User{
		UserName:    input.UserName,
		Password:    string(password),
		CreatedBy:   input.UserName,
		CreatedDate: time.Now(),
	}

	newUser, err := s.userRepository.Save(user)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *authService) Login(input input.LoginInput) (entity.User, error) {

	users, err := s.userRepository.FindByUserName(input.UserName)

	if err != nil {
		return entity.User{}, err
	}
	if len(users) == 0 {
		return entity.User{}, errors.New("user with username " + input.UserName + " not found")
	}

	user := users[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}
