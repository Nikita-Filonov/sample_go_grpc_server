package articles

import (
	articlesstore "github.com/Nikita-Filonov/sample_go_grpc_server/database/articles"
	articlesservice "github.com/Nikita-Filonov/sample_go_grpc_server/gen/proto"
	"github.com/Nikita-Filonov/sample_go_grpc_server/utils/logger"
	"google.golang.org/grpc"
)

type Server struct {
	articlesservice.UnimplementedArticlesServiceServer
	articlesStore articlesstore.Client
	logger        *logger.CtxLogger
}

func NewService(
	loggerService logger.Service,
	articlesStore articlesstore.Client,
) articlesservice.ArticlesServiceServer {
	return &Server{logger: loggerService.NewPrefix("API.ARTICLES"), articlesStore: articlesStore}
}

func RegisterService(
	grpcServer *grpc.Server,
	loggerService logger.Service,
	articlesStore articlesstore.Client,
) {
	articlesservice.RegisterArticlesServiceServer(grpcServer, NewService(loggerService, articlesStore))
}
