package todos

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetAllEndpoint endpoint.Endpoint
	GetEndpoint    endpoint.Endpoint
	CreateEndpoint endpoint.Endpoint
}

// ------------- Endpoints

func makeGetEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r := request.(getRequest)
		todo, err := srv.Get(ctx, r.Id)
		if err != nil {
			return response{Err: err.Error()}, err
		}

		return response{
			Id:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			Err:       "",
		}, nil
	}
}

func makeGetAllEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(getAllRequest)
		todos, err := srv.GetAll(ctx)
		if err != nil {
			return response{Err: err.Error()}, err
		}

		l := len(todos)
		res := make([]response, l, l)
		for i := 0; i < l; i++ {
			res[i] = response{
				Id:        todos[i].ID,
				Title:     todos[i].Title,
				Completed: todos[i].Completed,
				Err:       "",
			}
		}

		return res, nil
	}
}

func makeCreateEndpoint(srv Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r := request.(createRequest)

		todo, err := srv.Create(ctx, r)
		if err != nil {
			return response{Err: err.Error()}, err
		}

		return response{
			Id:        todo.ID,
			Title:     todo.Title,
			Completed: todo.Completed,
			Err:       "",
		}, nil
	}
}
