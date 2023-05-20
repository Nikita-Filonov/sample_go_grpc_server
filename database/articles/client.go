package articlesstore

import (
	"github.com/Nikita-Filonov/sample_go_grpc_server/database"
	"github.com/Nikita-Filonov/sample_go_grpc_server/utils/config"
	"github.com/Nikita-Filonov/sample_go_grpc_server/utils/logger"
)

type Client struct {
	database.Client
}

func NewClient(loggerService logger.Service, config config.Config) Client {
	client := database.NewClient(config, loggerService, "ARTICLES")

	return Client{client}
}
