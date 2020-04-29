package application

type UserID string

type User struct {
	ID        UserID
	Username  string
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

type Repository interface {
	Add(user User) error
	FindByID(id UserID) (User, error)
	Update(user User) error
	Delete(id UserID) error
}
