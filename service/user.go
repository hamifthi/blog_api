package service

import (
	"errors"
	"personal_blog/entity"
	"personal_blog/logger"
	"personal_blog/repository"
)

type UserService struct {
	userRepository repository.UserRepository
	logger         logger.Logger
}

func NewUserService(userRepository repository.UserRepository, logger logger.Logger) *UserService {
	return &UserService{userRepository: userRepository, logger: logger}
}

func (us *UserService) FindByID(id uint) (*entity.User, error) {
	user, err := us.userRepository.GetByID(id)
	if err != nil {
		us.logger.Error("facing error while finding user by ID: ", err, id)
		return nil, err
	}
	if user == nil {
		us.logger.Warn("couldn't find user by ID: ", id)
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (us *UserService) CreateUser(user *entity.User) error {
	if err := validateUser(user.Name, user.Email); err != nil {
		us.logger.Error("error in user validation: ", err, user)
		return err
	}
	_, err := us.userRepository.Create(user)
	if err != nil {
		us.logger.Error("couldn't create user: ", err, user)
		return err
	}
	return nil
}

func (us *UserService) UpdateUser(user *entity.User) error {
	if validateName(user.Name) == false {
		us.logger.Error("invalid user name: ", user)
		return errors.New("invalid user name")
	}

	_, err := us.FindByID(user.ID)
	if err != nil {
		return err
	}

	if err := us.userRepository.Update(user); err != nil {
		us.logger.Error("couldn't update user: ", err, user)
		return err
	}
	return nil
}

func (us *UserService) DeleteUser(id uint) error {
	user, err := us.FindByID(id)
	if err != nil {
		us.logger.Error("facing error while deleting user by ID: ", err, id)
		return err
	}
	return us.userRepository.Delete(user)
}
