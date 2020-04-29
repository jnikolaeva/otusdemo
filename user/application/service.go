package application

import "errors"

var ErrUserNotFound = errors.New("user not found")

type Service interface {
	Create(username, firstName, lastName, email, phone string) (UserID, error)
	FindByID(id string) (User, error)
	Update(id, firstName, lastName, email, phone string) error
	Delete(id string) error
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
		ID:        "111111", // @todo
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Add(user); err != nil {
		return "", err
	}

	return user.ID, nil
}

func (s service) FindByID(id string) (User, error) {
	user, err := s.repo.FindByID(UserID(id))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s service) Update(id, firstName, lastName, email, phone string) error {
	user, err := s.repo.FindByID(UserID(id))
	if err != nil {
		return err
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	user.Phone = phone

	return s.repo.Update(user)
}

func (s service) Delete(id string) error {
	return s.repo.Delete(UserID(id))
}
