package transport

import (
	"context"
	"github.com/arahna/otusdemo/user/application"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	FindUser   endpoint.Endpoint
	UpdateUser endpoint.Endpoint
	DeleteUser endpoint.Endpoint
}

func MakeEndpoints(s application.Service) Endpoints {
	return Endpoints{
		CreateUser: makeCreateUserEndpoint(s),
		FindUser:   makeFindUserEndpoint(s),
		UpdateUser: makeUpdateUserEndpoint(s),
		DeleteUser: makeDeleteUserEndpoint(s),
	}
}

func makeCreateUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		userID, err := s.Create(req.Username, req.FirstName, req.LastName, req.Phone, req.Email)
		if err != nil {
			return nil, err
		}
		return &createUserResponse{ID: string(userID)}, err
	}
}

func makeFindUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(findUserRequest)
		user, err := s.FindByID(req.ID)
		if err != nil {
			return nil, err
		}
		return &findUserResponse{
			userData: userData{
				ID: string(user.ID),
				userDetails: userDetails{
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					Phone:     user.Phone,
				},
			},
		}, err
	}
}

func makeUpdateUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateUserRequest)
		err := s.Update(req.ID, req.FirstName, req.LastName, req.Email, req.Phone)
		return &updateUserResponse{}, err
	}
}

func makeDeleteUserEndpoint(s application.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteUserRequest)
		err := s.Delete(req.ID)
		return &deleteUserResponse{}, err
	}
}
