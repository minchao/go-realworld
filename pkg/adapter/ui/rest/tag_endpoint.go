package rest

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/minchao/go-realworld/pkg/application/article"
	"github.com/minchao/go-realworld/pkg/application/article/domain"
)

type TagEndpoints struct {
	GetTags endpoint.Endpoint
}

func makeTagServerEndpoints(service article.TagUseCase) TagEndpoints {
	return TagEndpoints{
		GetTags: makeGetTagsEndpoint(service),
	}
}

type getTagsResponse struct {
	Tags []domain.Tag `json:"tags"`
}

func makeGetTagsEndpoint(service article.TagUseCase) endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (response interface{}, err error) {
		tags, err := service.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return getTagsResponse{Tags: tags}, nil
	}
}
