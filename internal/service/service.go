package service

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type Service struct {
	logger  *zap.Logger
	storage Storage
}

func New(logger *zap.Logger, storage Storage) *Service {
	return &Service{
		logger:  logger.Named("service"),
		storage: storage,
	}
}

func (s *Service) SaveData(ctx context.Context, key, data string) error {
	s.logger.Debug("saving data", zap.String("key", key), zap.String("data", data))

	if err := s.storage.Set(ctx, key, data); err != nil {
		return fmt.Errorf("set data in storage: %w", err)
	}
	return nil
}

func (s *Service) GetData(ctx context.Context, key string) (string, bool, error) {
	s.logger.Debug("getting data", zap.String("key", key))

	data, err := s.storage.Get(ctx, key)
	if err != nil {
		if errors.Is(err, ErrKeyNotFound) {
			return "", false, nil
		}
		return "", false, fmt.Errorf("get data in storage: %w", err)
	}

	return data, true, nil
}

func (s *Service) DeleteData(ctx context.Context, key string) error {
	s.logger.Debug("deleting data", zap.String("key", key))

	if err := s.storage.Delete(ctx, key); err != nil {
		return fmt.Errorf("delete data in storage: %w", err)
	}
	return nil
}
