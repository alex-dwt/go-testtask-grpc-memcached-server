package handlers

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/service"
	"github.com/alex-dwt/go-testtask-grpc-memcached-server/pkg/grpc_service"
)

type CacheHandler struct {
	cacheService *service.Service
	logger       *zap.Logger

	grpc_service.UnimplementedCacheServer
}

func NewCacheHandler(cacheService *service.Service, logger *zap.Logger) *CacheHandler {
	return &CacheHandler{
		logger:       logger.Named("grpc_server"),
		cacheService: cacheService,
	}
}

func (s *CacheHandler) Get(ctx context.Context, request *grpc_service.GetRequest) (*grpc_service.GetResponse, error) {
	data, found, err := s.cacheService.GetData(ctx, request.GetKey())
	if err != nil {
		s.logger.Error("failed to GetData", zap.Error(err))
		return nil, err
	}

	return &grpc_service.GetResponse{
		Value: data,
		Found: found,
	}, nil
}

func (s *CacheHandler) Set(ctx context.Context, request *grpc_service.SetRequest) (*emptypb.Empty, error) {
	if err := s.cacheService.SaveData(ctx, request.GetKey(), request.GetValue()); err != nil {
		s.logger.Error("failed to SaveData", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *CacheHandler) Delete(ctx context.Context, request *grpc_service.DeleteRequest) (*emptypb.Empty, error) {
	if err := s.cacheService.DeleteData(ctx, request.GetKey()); err != nil {
		s.logger.Error("failed to DeleteData", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
