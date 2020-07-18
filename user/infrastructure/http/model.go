package http

import "github.com/arahna/otusdemo/pkg/uuid"

type listUsersResponse struct {
	Users []*userData `json:"users"`
}

type createUserRequest struct {
	userDetails
}

type createUserResponse struct {
	ID string `json:"id"`
}

type findUserRequest struct {
	ID uuid.UUID `json:"userId"`
}

type findUserResponse struct {
	userData
}

type updateUserRequest struct {
	ID uuid.UUID `json:"userId"`
	userDetails
}

type updateUserResponse struct {
	userData
}

type deleteUserRequest struct {
	ID uuid.UUID `json:"userId"`
}

type errorResponse struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

type userData struct {
	ID string `json:"id"`
	userDetails
}

type userDetails struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
