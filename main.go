package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	articlesstore "sample_go_grpc_server/database/articles"
	"sample_go_grpc_server/server/articles"
	"sample_go_grpc_server/utils/config"
	"sample_go_grpc_server/utils/logger"
)

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))

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
