package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/arahna/otusdemo/pkg/uuid"
	"github.com/arahna/otusdemo/user/application"
	"github.com/go-kit/kit/log"
	gokittransport "github.com/go-kit/kit/transport"
	gokithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

func MakeHandler(pathPrefix string, endpoints Endpoints, errorLogger log.Logger, metrics *MetricsHolder) http.Handler {
	options := []gokithttp.ServerOption{
		gokithttp.ServerErrorEncoder(encodeErrorResponse),
		gokithttp.ServerErrorHandler(gokittransport.NewLogErrorHandler(errorLogger)),
	}

	listUsersHandler := gokithttp.NewServer(endpoints.ListUsers, decodeCreateUserRequest, encodeResponse, options...)
	createUserHandler := gokithttp.NewServer(endpoints.CreateUser, decodeCreateUserRequest, encodeResponse, options...)
	findUserHandler := gokithttp.NewServer(endpoints.FindUser, decodeFindUserRequest, encodeResponse, options...)
	updateUserHandler := gokithttp.NewServer(endpoints.UpdateUser, decodeUpdateUserRequest, encodeResponse, options...)
	deleteUserHandler := gokithttp.NewServer(endpoints.DeleteUser, decodeDeleteUserRequest, encodeResponse, options...)

	r := mux.NewRouter()
	s := r.PathPrefix(pathPrefix).Subrouter()
	s.Handle("", instrumentingMiddleware(listUsersHandler, metrics, "ListUsers")).Methods(http.MethodGet)
	s.Handle("", instrumentingMiddleware(createUserHandler, metrics, "CreateUser")).Methods(http.MethodPost)
	s.Handle("/{userId}", instrumentingMiddleware(findUserHandler, metrics, "FindUser")).Methods(http.MethodGet)
	s.Handle("/{userId}", instrumentingMiddleware(updateUserHandler, metrics, "UpdateUser")).Methods(http.MethodPut)
	s.Handle("/{userId}", instrumentingMiddleware(deleteUserHandler, metrics, "DeleteUser")).Methods(http.MethodDelete)
	return r
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req createUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil && e != io.EOF {
		return nil, e
	}
	return req, nil
}

func decodeFindUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	sID, ok := vars["userId"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := uuid.FromString(sID)
	if err != nil {
		return nil, ErrBadRouting
	}
	return findUserRequest{ID: id}, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	sID, ok := vars["userId"]
	if !ok {
		return nil, ErrBadRouting
	}
	var req updateUserRequest
	if e := json.NewDecoder(r.Body).Decode(&req.userDetails); e != nil && e != io.EOF {
		return nil, e
	}
	id, err := uuid.FromString(sID)
	if err != nil {
		return nil, ErrBadRouting
	}
	req.ID = id
	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	sID, ok := vars["userId"]
	if !ok {
		return nil, ErrBadRouting
	}
	id, err := uuid.FromString(sID)
	if err != nil {
		return nil, ErrBadRouting
	}
	return deleteUserRequest{ID: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var errorResponse = translateError(err)
	w.WriteHeader(errorResponse.Status)
	_ = json.NewEncoder(w).Encode(errorResponse.Response)
}

type transportError struct {
	Status   int
	Response errorResponse
}

func translateError(err error) transportError {
	switch err {
	case application.ErrUserNotFound:
		return transportError{
			Status: http.StatusNotFound,
			Response: errorResponse{
				Code:    101,
				Message: err.Error(),
			},
		}
	case application.ErrDuplicateUser:
		return transportError{
			Status: http.StatusConflict,
			Response: errorResponse{
				Code:    102,
				Message: err.Error(),
			},
		}
	default:
		return transportError{
			Status: http.StatusInternalServerError,
			Response: errorResponse{
				Code:    100,
				Message: "unexpected error",
			},
		}
	}
}
