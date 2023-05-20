package articles

import (
	"google.golang.org/grpc"
	articlesstore "sample_go_grpc_server/database/articles"
	articlesservice "sample_go_grpc_server/gen/proto"
	"sample_go_grpc_server/utils/logger"
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
