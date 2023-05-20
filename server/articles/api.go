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

	return &articlesservice.GetArticleResponse{Article: mappings.MapArticleModelToApiArticle(article)}, err
}

func (s *Server) CreateArticle(
	_ context.Context,
	request *articlesservice.CreateArticleRequest,
) (*articlesservice.CreateArticleResponse, error) {
	s.logger.InfofJSON("CreateArticleRequest", request)

	article, err := s.articlesStore.CreateArticle(request.Article)

	return &articlesservice.CreateArticleResponse{Article: mappings.MapArticleModelToApiArticle(article)}, err
}

func (s *Server) UpdateArticle(
	_ context.Context,
	request *articlesservice.UpdateArticleRequest,
) (*articlesservice.UpdateArticleResponse, error) {
	s.logger.InfofJSON("UpdateArticleRequest", request)

	article, err := s.articlesStore.UpdateArticle(request.Article)

	return &articlesservice.UpdateArticleResponse{Article: mappings.MapArticleModelToApiArticle(article)}, err
}

func (s *Server) DeleteArticle(
	_ context.Context,
	request *articlesservice.DeleteArticleRequest,
) (*articlesservice.DeleteArticleResponse, error) {
	s.logger.InfofJSON("DeleteArticleRequest", request)

	err := s.articlesStore.DeleteArticle(request.ArticleId)

	return &articlesservice.DeleteArticleResponse{}, err
}
