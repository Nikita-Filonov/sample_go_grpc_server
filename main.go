package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	articlesstore "github.com/Nikita-Filonov/sample_go_grpc_server/database/articles"
	"github.com/Nikita-Filonov/sample_go_grpc_server/server/articles"
	"github.com/Nikita-Filonov/sample_go_grpc_server/utils/config"
	"github.com/Nikita-Filonov/sample_go_grpc_server/utils/logger"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	configuration := config.NewConfig()
	loggerService := logger.NewLoggerService(configuration)

	articleStore := articlesstore.NewClient(loggerService, configuration)

	articles.RegisterService(grpcServer, loggerService, articleStore)

	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
