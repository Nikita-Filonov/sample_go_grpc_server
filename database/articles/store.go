package articlesstore

import (
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
)

type Store interface {
	GetArticle(articleId string) (ArticleModel, error)
	CreateArticle(payload *articlesservice.Article) (ArticleModel, error)
	UpdateArticle(payload *articlesservice.Article) (ArticleModel, error)
	DeleteArticle(articleId string) error
}

const (
	getArticleQuery    = "SELECT id, title, author, description FROM articles WHERE id=?"
	createArticleQuery = "INSERT INTO articles (id, title, author, description) VALUES (?, ?, ?, ?) RETURNING *"
	updateArticleQuery = "UPDATE articles SET title=?, author=?, description=? WHERE id=? RETURNING *"
	deleteArticleQuery = "DELETE FROM articles WHERE id=?"
)

func (c *Client) GetArticle(articleId string) (ArticleModel, error) {
	var model = ArticleModel{}
	row := c.QueryRow(getArticleQuery, articleId)

	err := row.Scan(&model.Id, &model.Title, &model.Author, &model.Description)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (c *Client) CreateArticle(payload *articlesservice.Article) (ArticleModel, error) {
	var model = ArticleModel{}
	row := c.QueryRow(createArticleQuery, payload.Id, payload.Title, payload.Author, payload.Description)

	err := row.Scan(&model.Id, &model.Title, &model.Author, &model.Description)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (c *Client) UpdateArticle(payload *articlesservice.Article) (ArticleModel, error) {
	var model = ArticleModel{}
	row := c.QueryRow(updateArticleQuery, payload.Title, payload.Author, payload.Description, payload.Id)

	err := row.Scan(&model.Id, &model.Title, &model.Author, &model.Description)
	if err != nil {
		return model, err
	}

	return model, nil
}

func (c *Client) DeleteArticle(articleId string) error {
	_, err := c.Exec(deleteArticleQuery, articleId)

	return err
}
