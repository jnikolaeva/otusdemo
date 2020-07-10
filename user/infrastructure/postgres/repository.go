package postgres

import (
	"github.com/arahna/otusdemo/pkg/uuid"
	"github.com/arahna/otusdemo/user/application"
	"github.com/jackc/pgx"
	"github.com/pkg/errors"
)

const errUniqueConstraint = "23505"

type rawUser struct {
	ID        string `db:"id"`
	Username  string `db:"username"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
}

type repository struct {
	connPool *pgx.ConnPool
}

func New(connPool *pgx.ConnPool) application.Repository {
	return &repository{
		connPool: connPool,
	}
}

func (r *repository) NextID() application.UserID {
	return application.UserID(uuid.Generate())
}

func (r *repository) Add(user application.User) error {
	_, err := r.connPool.Exec(
		"INSERT INTO users (id, username, first_name, last_name, email, phone) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID.String(), user.Username, user.FirstName, user.LastName, user.Email, user.Phone)
	return r.convertError(err)
}

func (r *repository) FindByID(id application.UserID) (application.User, error) {
	var raw rawUser
	query := "SELECT id, username, first_name, last_name, phone, email FROM users WHERE id = $1"
	err := r.connPool.QueryRow(query, id.String()).Scan(&raw.ID, &raw.Username, &raw.FirstName, &raw.LastName, &raw.Phone, &raw.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = application.ErrUserNotFound
		}
		return application.User{}, errors.WithStack(err)
	}
	userID, _ := uuid.FromString(raw.ID)
	user := application.User{
		ID:        application.UserID(userID),
		Username:  raw.Username,
		FirstName: raw.FirstName,
		LastName:  raw.LastName,
		Email:     raw.Email,
		Phone:     raw.Phone,
	}
	return user, nil
}

func (r *repository) Find() ([]*application.User, error) {
	var users []*application.User
	query := "SELECT id, username, first_name, last_name, phone, email FROM users"
	rows, err := r.connPool.Query(query)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = application.ErrUserNotFound
		}
		return users, errors.WithStack(err)
	}
	defer rows.Close()

	var raw rawUser
	for rows.Next() {
		var user application.User
		if err = rows.Scan(&raw.ID, &raw.Username, &raw.FirstName, &raw.LastName, &raw.Phone, &raw.Email); err != nil {
			return users, err
		}
		userID, _ := uuid.FromString(raw.ID)
		user = application.User{
			ID:        application.UserID(userID),
			Username:  raw.Username,
			FirstName: raw.FirstName,
			LastName:  raw.LastName,
			Email:     raw.Email,
			Phone:     raw.Phone,
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *repository) Update(user application.User) error {
	_, err := r.connPool.Exec(
		"UPDATE users SET username = $1, first_name = $2, last_name = $3, email = $4, phone = $5 WHERE id = $6",
		user.Username, user.FirstName, user.LastName, user.Email, user.Phone, user.ID.String())
	return r.convertError(err)
}

func (r *repository) Delete(id application.UserID) error {
	_, err := r.connPool.Exec("DELETE FROM users WHERE id = $1", id.String())
	return errors.WithStack(err)
}

func (r *repository) convertError(err error) error {
	if err != nil {
		pgErr, ok := err.(pgx.PgError)
		if ok && pgErr.Code == errUniqueConstraint {
			return application.ErrDuplicateUser
		}
		return errors.WithStack(err)
	}
	return nil
}
