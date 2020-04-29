package transport

type createUserRequest struct {
	userDetails
}

type createUserResponse struct {
	ID string `json:"id"`
}

type findUserRequest struct {
	ID string `json:"userId"`
}

type findUserResponse struct {
	userData
}

type updateUserRequest struct {
	ID string `json:"userId"`
	userDetails
}

type updateUserResponse struct {
}

type deleteUserRequest struct {
	ID string `json:"userId"`
}

type deleteUserResponse struct {
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
