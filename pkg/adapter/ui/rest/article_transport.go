package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/minchao/go-realworld/pkg/application/article"
)

func InitArticleHandler(router *mux.Router, service article.UseCase, options []httptransport.ServerOption) {
	endpoints := makeArticleServerEndpoints(service)

	router.Methods("GET").Path("/articles").Handler(httptransport.NewServer(
		endpoints.GetArticles,
		BypassRequest,
		EncodeResponse,
		options...,
	))

	router.Methods("GET").Path("/articles/{slug}").Handler(httptransport.NewServer(
		endpoints.GetArticle,
		decodeGetArticleRequest,
		EncodeResponse,
		options...,
	))

	router.Methods("POST").Path("/articles").Handler(httptransport.NewServer(
		endpoints.PostArticle,
		decodePostArticlesRequest,
		EncodeResponse,
		options...,
	))
}

func decodeGetArticleRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]
	if !ok {
		return nil, errors.New("slug not found")
	}
	return getArticleRequest{Slug: slug}, nil
}

func decodePostArticlesRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req postArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
