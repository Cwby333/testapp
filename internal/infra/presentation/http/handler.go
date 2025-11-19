package http

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/Cwby333/testapp/pkg/api/v1"
)

type Repo interface {
	Insert(ctx context.Context, number int) error
	Get(ctx context.Context) ([]int, error)
}

type Handler struct {
	mux *http.ServeMux
	repo Repo
}

func New(repo Repo) *Handler {
	return &Handler{
		mux: http.NewServeMux(),
		repo: repo,
	}
}

func (h Handler) RegisterRoutes(strict api.ServerInterface) {
	h.mux.Handle("POST /number", http.HandlerFunc(strict.NumberPost))
	h.mux.Handle("GET /number", http.HandlerFunc(strict.NumberGet))

	h.mux.Handle("GET /hello", http.HandlerFunc(strict.HelloGet))
}

func (h *Handler) HelloGet(ctx context.Context, request api.HelloGetRequestObject) (api.HelloGetResponseObject, error) {
	s := "Hello from test app"
	return api.HelloGet200JSONResponse{Answer: &s}, nil
}

func (h *Handler) NumberGet(ctx context.Context, request api.NumberGetRequestObject) (api.NumberGetResponseObject, error) {
	numbers, err := h.repo.Get(ctx)
	if err != nil {
		return api.NumberGet500Response{}, err
	}

	slices.Sort(numbers)

	return api.NumberGet200JSONResponse{
		Numbers: &numbers,
	}, nil
}

func (h *Handler) NumberPost(ctx context.Context, request api.NumberPostRequestObject) (api.NumberPostResponseObject, error) {
	number := request.Body.Number
	fmt.Println(number)
	err := h.repo.Insert(ctx, *number)
	if err != nil {
		return api.NumberPost500Response{}, nil
	}

	numbers, err := h.repo.Get(ctx)
	if err != nil {
		return api.NumberPost500Response{}, err
	}

	slices.Sort(numbers)

	return api.NumberPost200JSONResponse{Numbers: &numbers}, nil
}