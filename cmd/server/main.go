package main

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/adapters/handlers"
	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/adapters/storage/memory"
	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/service"
	"github.com/alex-dwt/go-testtask-grpc-memcached-server/pkg/grpc_service"
)

func main() {
	//logger, err := zap.NewProduction()
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}
	defer logger.Sync()

	storage := memory.New()
	cacheService := service.New(logger, storage)

	logger.Info("starting server...")
	lis, err := net.Listen("tcp", ":8889")
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}

	server := grpc.NewServer()
	grpc_service.RegisterCacheServer(server, handlers.NewCacheHandler(cacheService, logger))
	if err := server.Serve(lis); err != nil {
		logger.Fatal("failed to serve", zap.Error(err))
	}
}
