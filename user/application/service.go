package application

import (
	"errors"
	"github.com/arahna/otusdemo/pkg/uuid"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrDuplicateUser = errors.New("user with such username already exists")
)

type Service interface {
	Create(username, firstName, lastName, email, phone string) (UserID, error)
	FindByID(id uuid.UUID) (User, error)
	List() ([]*User, error)
	Update(id uuid.UUID, username, firstName, lastName, email, phone string) (User, error)
	Delete(id uuid.UUID) error
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

type service struct {
	repo Repository
}

func (s service) Create(username, firstName, lastName, email, phone string) (UserID, error) {
	user := User{
		ID:        s.repo.NextID(),
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Add(user); err != nil {
		return user.ID, err
	}

	return user.ID, nil
}

func (s service) FindByID(id uuid.UUID) (User, error) {
	return s.repo.FindByID(UserID(id))
}

func (s service) List() ([]*User, error) {
	return s.repo.Find()
}

func (s service) Update(id uuid.UUID, username, firstName, lastName, email, phone string) (User, error) {
	user, err := s.repo.FindByID(UserID(id))
	if err != nil {
		return user, err
	}

	user.Username = username
	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.Phone = phone

	return user, s.repo.Update(user)
}

func (s service) Delete(id uuid.UUID) error {
	return s.repo.Delete(UserID(id))
}
