package mappings

import (
	articlesstore "sample_go_grpc_server/database/articles"
	articlesservice "sample_go_grpc_server/gen/proto"
)

func MapArticleModelToApiArticle(model articlesstore.ArticleModel) *articlesservice.Article {
	return &articlesservice.Article{
		Id:          model.Id,
		Title:       model.Title,
		Author:      model.Author,
		Description: model.Description,
	}
}
