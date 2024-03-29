package service

import (
	"errors"
	"math/rand"
	"myfinnplan/entity"
	"myfinnplan/helper"
	"myfinnplan/input"
	"myfinnplan/repository"
	"net/smtp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(input input.UserInput) (entity.User, error)
	Login(input input.LoginInput) (entity.User, error)
	VerifiedUser(id int) (entity.User, error)
	SendEmail(user entity.User, token string) error
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) *authService {
	return &authService{userRepository}
}

func (s *authService) SendEmail(user entity.User, token string) error {
	email := entity.SetEnv().EMAIL
	pw := entity.SetEnv().EMAIL_PASS
	host := "smtp.gmail.com"

	auth := smtp.PlainAuth("", email, pw, host)

	request := helper.NewRequest([]string{user.Email}, "[No-Reply]", "Send")

	a := helper.Replace{
		ID:    user.Id,
		Link:  entity.SetEnv().URL,
		Token: token,
		Name:  user.UserName,
	}

	err := request.ParseTemplate("email.html", a)
	if err != nil {
		return err
	}

	ok, err := request.SendEmail(auth)
	if !ok {
		return err
	}

	// err = smtp.SendMail(host+":"+port, auth, email, []string{user.Email}, []byte(msg))

	if err != nil {
		return err
	}
	return nil
}

func (s *authService) VerifiedUser(id int) (entity.User, error) {
	checkUser, err := s.userRepository.FindById(id)

	if err != nil {
		return entity.User{}, errors.New("error find user")
	}

	if len(checkUser) == 0 {
		return entity.User{}, errors.New("user not found")
	}

	userOld := checkUser[0]

	userOld.IsVerified = true
	userOld.UpdatedBy = userOld.UserName
	userOld.UpdatedDate = time.Now()

	user, err := s.userRepository.Edit(userOld)

	if err != nil {
		return entity.User{}, errors.New("error save user")
	}

	return user, nil

}

func (s *authService) RegisterUser(input input.UserInput) (entity.User, error) {

	checkUser, err := s.userRepository.FindByUserName(input.UserName)

	if err != nil {
		return entity.User{}, errors.New("error find user")
	}

	if len(checkUser) != 0 {
		return entity.User{}, errors.New("UserName already used")
	}

	checkUser, err = s.userRepository.FindByEmail(input.Email)

	if err != nil {
		return entity.User{}, errors.New("error find user")
	}

	if len(checkUser) != 0 {
		return entity.User{}, errors.New("email already used")
	}

	if len(strings.Split(input.Email, "@")) != 2 || len(strings.Split(strings.Split(input.Email, "@")[1], ".")) < 2 {
		return entity.User{}, errors.New("invalid email")
	}

	key := rand.Intn(9)
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), key)
	if err != nil {
		return entity.User{}, errors.New("error encrypt password")
	}

	user := entity.User{
		UserName:    input.UserName,
		Password:    string(password),
		Email:       input.Email,
		Photo:       input.Photo,
		Telephone:   input.Telephone,
		IsVerified:  false,
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
		users, err = s.userRepository.FindByEmail(input.UserName)
		if err != nil {
			return entity.User{}, err
		}
		if len(users) == 0 {
			return entity.User{}, errors.New("user with username or email " + input.UserName + " not found")
		}
	}

	user := users[0]

	if !user.IsVerified {
		return user, errors.New("email not verified yet")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}
