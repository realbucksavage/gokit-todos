package todos

import (
	"context"

	"github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"gokit-todos/database"
)

func AddRoutes(ctx context.Context, r *mux.Router) {
	service := NewService(ctx.Value(database.ContextKey).(*gorm.DB))

	endpoints := Endpoints{
		GetEndpoint:    makeGetEndpoint(service),
		GetAllEndpoint: makeGetAllEndpoint(service),
		CreateEndpoint: makeCreateEndpoint(service),
	}

	r.Methods("GET").Path("/").Handler(http.NewServer(
		endpoints.GetAllEndpoint,
		decodeGetAllRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/{id}").Handler(http.NewServer(
		endpoints.GetEndpoint,
		decodeGetRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/").Handler(http.NewServer(
		endpoints.CreateEndpoint,
		decodeCreateRequest,
		encodeResponse,
	))
}
