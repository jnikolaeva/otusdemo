package application

import "github.com/arahna/otusdemo/pkg/uuid"

type UserID uuid.UUID

func (u UserID) String() string {
	return uuid.UUID(u).String()
}

type User struct {
	ID        UserID
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Repository interface {
	NextID() UserID
	Add(user User) error
	FindByID(id UserID) (User, error)
	Update(user User) error
	Delete(id UserID) error
}
