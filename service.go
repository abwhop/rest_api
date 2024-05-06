package rest_api

import (
	"context"
	"errors"
)

// ServiceInterface Base interface of base service type
type ServiceInterface interface {
	GetList(context.Context, func(interface{}) error, func(user interface{}) error) (interface{}, ErrorInterface)
	GetById(context.Context, string, func(interface{}) error, func(user interface{}) error) (interface{}, ErrorInterface)
	Create(context.Context, func(interface{}) error, func(interface{}) error, func(user interface{}) error) (interface{}, ErrorInterface)
	Update(context.Context, string, func(interface{}) error, func(interface{}) error, func(user interface{}) error) (interface{}, ErrorInterface)
	Delete(context.Context, string, func(interface{}) error, func(user interface{}) error) (interface{}, ErrorInterface)
}

// Service Base service type
type Service struct{}

// GetList Method for getting a list. For overriding
func (s *Service) GetList(ctx context.Context, getQuery func(interface{}) error, getUser func(interface{}) error) (interface{}, ErrorInterface) {
	return nil, NewNotImplementedError(errors.New("method not implemented"))
}

// GetById Method for getting. For overriding
func (s *Service) GetById(ctx context.Context, id string, getQuery func(interface{}) error, getUser func(interface{}) error) (interface{}, ErrorInterface) {
	return nil, NewNotImplementedError(errors.New("method not implemented"))
}

// Create Method for create. For overriding
func (s *Service) Create(ctx context.Context, getQuery func(interface{}) error, getData func(interface{}) error, getUser func(interface{}) error) (interface{}, ErrorInterface) {
	return nil, NewNotImplementedError(errors.New("method not implemented"))
}

// Update Method for update. For overriding
func (s *Service) Update(ctx context.Context, id string, getQuery func(interface{}) error, getData func(interface{}) error, getUser func(interface{}) error) (interface{}, ErrorInterface) {
	return nil, NewNotImplementedError(errors.New("method not implemented"))
}

// Delete Method for delete. For overriding
func (s *Service) Delete(ctx context.Context, id string, binding func(interface{}) error, getUser func(interface{}) error) (interface{}, ErrorInterface) {
	return nil, NewNotImplementedError(errors.New("method not implemented"))
}
