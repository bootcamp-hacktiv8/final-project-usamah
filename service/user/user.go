package user

import (
	"errors"
	"final-project-usamah/delivery/helper"
	_request "final-project-usamah/delivery/helper/request/user"
	_entities "final-project-usamah/entities"
	_userRepository "final-project-usamah/repository/user"
	"strings"
	"time"
)

type UserService struct {
	userRepository _userRepository.UserRepositoryInterface
}

func NewUserService(userRepository _userRepository.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepository: userRepository,
	}
}

func (us *UserService) Register(user _entities.User) (_entities.User, error) {
	//validasi saat registrasi
	if user.Username == "" {
		return user, errors.New("name is required")
	}
	if user.Email == "" {
		return user, errors.New("email is required")
	}
	if !strings.Contains(user.Email, "@") {
		return user, errors.New("wrong email format")
	}
	if user.Password == "" {
		return user, errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return user, errors.New("minimum number of password characters is 6")
	}
	if user.Age == 0 {
		return user, errors.New("age is required")
	}
	if user.Age <= 8 {
		return user, errors.New("age must be above 8 years")
	}

	password, _ := helper.HashPassword(user.Password)
	user.Password = password

	newUser, id, err := us.userRepository.Register(user)
	newUser.Id = id
	return newUser, err
}

func (us *UserService) Login(inputLogin _request.FormatLogin) (_entities.User, error) {

	user, err := us.userRepository.Login(inputLogin.Email)

	if user.Email == "" {
		return user, errors.New("email incorrect")
	}

	errCheck := helper.CheckPassHash(inputLogin.Password, user.Password)
	if errCheck != nil {
		return user, errors.New("password incorrect")
	}

	return user, err
}

func (us *UserService) GetUser(idToken int) (_entities.User, error) {
	user, err := us.userRepository.GetUser(idToken)
	return user, err
}

func (us *UserService) UpdateUser(updateUser _entities.User, idToken int) (_entities.User, error) {
	getUser, err := us.userRepository.GetUser(idToken)
	if err != nil {
		return getUser, err
	}

	//validasi saat update user
	if updateUser.Username != "" {
		getUser.Username = updateUser.Username
	}
	if updateUser.Email != "" {
		if !strings.Contains(updateUser.Email, "@") {
			return updateUser, errors.New("wrong email format")
		}
		getUser.Email = updateUser.Email
	}
	if updateUser.Age != 0 {
		if updateUser.Age <= 8 {
			return updateUser, errors.New("age must be above 8 years")
		}
		getUser.Age = updateUser.Age
	}

	user, err := us.userRepository.UpdateUser(getUser, idToken)
	user.Id = idToken
	user.Updated_at = time.Now()
	return user, err
}

func (us *UserService) DeleteUser(idToken int) error {
	err := us.userRepository.DeleteUser(idToken)
	return err
}