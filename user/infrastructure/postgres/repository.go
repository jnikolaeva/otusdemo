package postgres

import (
	"github.com/arahna/otusdemo/user/application"
)

type repository struct{}

func New() application.Repository {
	return repository{}
}

func (repository) Add(user application.User) error {
	return nil
}

func (repository) FindByID(id application.UserID) (application.User, error) {
	return application.User{
		ID:        id,
		Username:  "username1",
		FirstName: "First Name",
		Email:     "email@domain.com",
	}, nil
}

func (repository) Update(user application.User) error {
	return nil
}

func (repository) Delete(id application.UserID) error {
	return nil
}
