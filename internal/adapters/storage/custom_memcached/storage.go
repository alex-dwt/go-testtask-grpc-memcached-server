package custom_memcached

import (
	"context"
	"fmt"

	"github.com/alex-dwt/go-testtask-grpc-memcached-cache/pkg/client"

	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/service"
)

type Storage struct {
	client *client.Client
}

func New(addr string) (*Storage, error) {
	c, err := client.New(addr)
	if err != nil {
		return nil, fmt.Errorf("client create: %w", err)
	}

	return &Storage{
		client: c,
	}, nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	val, found, err := s.client.Get(ctx, key)
	if err != nil {
		return "", fmt.Errorf("client: %w", err)
	}

	if !found {
		return "", service.ErrKeyNotFound
	}

	return val, nil
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	if err := s.client.Set(ctx, key, value); err != nil {
		return fmt.Errorf("client: %w", err)
	}

	return nil
}

func (s *Storage) Delete(ctx context.Context, key string) error {
	if err := s.client.Delete(ctx, key); err != nil {
		return fmt.Errorf("client: %w", err)
	}

	return nil
}
