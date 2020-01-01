package rest

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/minchao/go-realworld/pkg/application/article"
)

func InitTagHandler(router *mux.Router, service article.TagUseCase, options []httptransport.ServerOption) {
	endpoints := makeTagServerEndpoints(service)

	router.Methods("GET").Path("/tags").Handler(httptransport.NewServer(
		endpoints.GetTags,
		BypassRequest,
		EncodeResponse,
		options...,
	))
}
