package articles

import (
	"context"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/Nikita-Filonov/sample_go_grpc_server/utils/mappings"
)

func (s *Server) GetArticle(
	_ context.Context,
	request *articlesservice.GetArticleRequest,
) (*articlesservice.GetArticleResponse, error) {
	s.logger.InfofJSON("GetArticleRequest", request)

	article, err := s.articlesStore.GetArticle(request.ArticleId)

	if err != nil {
		return mappings.GetArticleNotFoundError(request.ArticleId), nil
	}

	return &articlesservice.GetArticleResponse{
		Result: &articlesservice.GetArticleResponse_Article{
			Article: mappings.MapArticleModelToResponse(article),
		},
	}, nil
}

func (s *Server) CreateArticle(
	_ context.Context,
	request *articlesservice.CreateArticleRequest,
) (*articlesservice.CreateArticleResponse, error) {
	s.logger.InfofJSON("CreateArticleRequest", request)

	article, err := s.articlesStore.CreateArticle(request.Article)

	return &articlesservice.CreateArticleResponse{
		Result: &articlesservice.CreateArticleResponse_Article{
			Article: mappings.MapArticleModelToResponse(article),
		},
	}, err
}

func (s *Server) UpdateArticle(
	_ context.Context,
	request *articlesservice.UpdateArticleRequest,
) (*articlesservice.UpdateArticleResponse, error) {
	s.logger.InfofJSON("UpdateArticleRequest", request)

	article, err := s.articlesStore.UpdateArticle(request.Article)

	return &articlesservice.UpdateArticleResponse{
		Result: &articlesservice.UpdateArticleResponse_Article{
			Article: mappings.MapArticleModelToResponse(article),
		},
	}, err
}

func (s *Server) DeleteArticle(
	_ context.Context,
	request *articlesservice.DeleteArticleRequest,
) (*articlesservice.DeleteArticleResponse, error) {
	s.logger.InfofJSON("DeleteArticleRequest", request)

	err := s.articlesStore.DeleteArticle(request.ArticleId)

	return &articlesservice.DeleteArticleResponse{}, err
}
