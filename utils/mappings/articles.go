package mappings

import (
	"fmt"
	articlesstore "github.com/Nikita-Filonov/sample_go_grpc_server/database/articles"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
)

func MapArticleModelToResponse(model articlesstore.ArticleModel) *articlesservice.Article {
	return &articlesservice.Article{
		Id:          model.Id,
		Title:       model.Title,
		Author:      model.Author,
		Description: model.Description,
	}
}

func GetArticleNotFoundError(articleId string) *articlesservice.GetArticleResponse {
	return &articlesservice.GetArticleResponse{
		Result: &articlesservice.GetArticleResponse_Error{
			Error: &articlesservice.Error{
				Type:    articlesservice.ErrorType_NOT_FOUND,
				Message: fmt.Sprintf("Article with Id %s not found", articleId),
			},
		},
	}
}
