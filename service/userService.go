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

type UserService interface {
	CreateUser(input input.UserInput, userName string) (entity.User, error)
	EditUser(input input.UserEditInput, userName string) (entity.User, error)
	GetUserById(id int) ([]entity.User, error)
	GetUserByUserName(userName string) ([]entity.User, error)

	GetAllUser() ([]entity.User, error)
	DeleteUser(id int, userName string) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input input.UserInput, userName string) (entity.User, error) {
	user := entity.User{
		UserName:    input.UserName,
		Password:    input.Password,
		CreatedBy:   userName,
		CreatedDate: time.Now(),
	}

	newUser, err := s.userRepository.Save(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (s *userService) EditUser(input input.UserEditInput, userName string) (entity.User, error) {
	oldUsers, err := s.userRepository.FindById(input.Id)

	if err != nil {
		return entity.User{}, err
	}
	if input.UserName != oldUsers[0].UserName {
		checkUser, err := s.userRepository.FindByUserName(input.UserName)

		if err != nil {
			return entity.User{}, errors.New("error find user")
		}

		if len(checkUser) != 0 {
			return entity.User{}, errors.New("UserName already used")
		}
	}

	key := rand.Intn(9)
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), key)
	if err != nil {
		return entity.User{}, errors.New("error encrypt password")
	}

	oldUser := oldUsers[0]

	user := entity.User{
		Id:          input.Id,
		UserName:    input.UserName,
		Password:    string(password),
		Telephone:   input.Telephone,
		Photo:       input.Photo,
		CreatedBy:   oldUser.CreatedBy,
		CreatedDate: oldUser.CreatedDate,
		UpdatedBy:   userName,
		UpdatedDate: time.Now(),
	}

	newUser, err := s.userRepository.Edit(user)

	if err != nil {
		return user, err
	}

	return newUser, nil
}

func (s *userService) GetUserById(id int) ([]entity.User, error) {

	user, err := s.userRepository.FindById(id)

	if err != nil {
		return user, err
	}

	if len(user) == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}
func (s *userService) GetUserByUserName(userName string) ([]entity.User, error) {

	user, err := s.userRepository.FindByUserName(userName)

	if err != nil {
		return user, err
	}

	if len(user) == 0 {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (s *userService) GetAllUser() ([]entity.User, error) {
	users, err := s.userRepository.FindAll()

	if err != nil {
		return users, err
	}

	if len(users) <= 0 {
		return users, errors.New("user not found")
	}

	return users, nil
}

func (s *userService) DeleteUser(id int, userName string) (entity.User, error) {
	users, err := s.GetUserById(id)

	if err != nil {
		return entity.User{}, err
	}

	if len(users) == 0 {
		return entity.User{}, errors.New("user not found")
	}

	user := users[0]

	user.DeletedDate = time.Now()
	user.DeletedBy = userName
	result, err := s.userRepository.Edit(user)
	if err != nil {
		return result, err
	}
	return result, nil
}
