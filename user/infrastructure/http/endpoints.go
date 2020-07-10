package http

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/arahna/otusdemo/user/application"
)

type Endpoints struct {
	ListUsers  endpoint.Endpoint
	CreateUser endpoint.Endpoint
	FindUser   endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	DeleteUser endpoint.Endpoint
}

func MakeEndpoints(s application.Service) Endpoints {
	return Endpoints{
		ListUsers:  makeListUsersEndpoint(s),
		CreateUser: makeCreateUserEndpoint(s),
		FindUser:   makeFindUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
	}
}

func makeListUsersEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.List()
		if err != nil {
			return &listUsersResponse{}, err
		}
		var userList []*userData
		for _, user := range users {
			userData := toUserData(*user)
			userList = append(userList, &userData)
		}
		return &listUsersResponse{userList}, nil
	}
}

func makeCreateUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		userID, err := s.Create(req.Username, req.FirstName, req.LastName, req.Phone, req.Email)
		if err != nil {
			return nil, err
		}
		return &createUserResponse{ID: userID.String()}, err
	}
}

func makeFindUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findUserRequest)
		user, err := s.FindByID(req.ID)
		if err != nil {
			return nil, err
		}
		return &findUserResponse{toUserData(user)}, err
	}
}

func makeUpdateUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		user, err := s.Update(req.ID, req.Username, req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			return nil, err
		}
		return &updateUserResponse{toUserData(user)}, err
	}
}

func makeDeleteUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		return nil, s.Delete(req.ID)
	}
}

func toUserData(user application.User) userData {
	return userData{
		ID: user.ID.String(),
		userDetails: userDetails{
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		},
	}
}
