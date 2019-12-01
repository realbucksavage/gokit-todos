package todos

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type getRequest struct {
	Id int
}

type getAllRequest struct{}

type createRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type response struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`

	Err string `json:"err,omitempty"`
}

// ---------- Decoder methods

func decodeGetRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return getRequest{Id: id}, nil
}

func decodeGetAllRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req getAllRequest
	return req, nil
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}

// -------- Encode

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
