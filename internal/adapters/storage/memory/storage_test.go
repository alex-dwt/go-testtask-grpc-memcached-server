package memory_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/adapters/storage/memory"
	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/service"
)

func TestStorage_Get_OK(t *testing.T) {
	const (
		k = "some-key"
		v = "some-value"
	)
	ctx := context.Background()

	storage := memory.New()
	err := storage.Set(ctx, k, v)
	require.NoError(t, err)

	value, err := storage.Get(ctx, k)
	assert.Equal(t, v, value)
	assert.NoError(t, err)
}

func TestStorage_Get_NotFoundErr(t *testing.T) {
	storage := memory.New()
	_, err := storage.Get(context.Background(), "k")
	assert.ErrorIs(t, err, service.ErrKeyNotFound)
}

func TestStorage_Delete(t *testing.T) {
	const (
		k = "some-key"
		v = "some-value"
	)
	ctx := context.Background()

	storage := memory.New()
	err := storage.Set(ctx, k, v)
	require.NoError(t, err)

	value, err := storage.Get(ctx, k)
	require.Equal(t, v, value)
	require.NoError(t, err)

	err = storage.Delete(ctx, k)
	require.NoError(t, err)

	_, err = storage.Get(ctx, k)
	assert.ErrorIs(t, err, service.ErrKeyNotFound)
}

func TestStorage_Set(t *testing.T) {
	const (
		k = "some-key"
		v = "some-value"
	)
	ctx := context.Background()

	storage := memory.New()
	err := storage.Set(ctx, k, v)
	require.NoError(t, err)

	value, err := storage.Get(ctx, k)
	require.NoError(t, err)
	assert.Equal(t, v, value, "Value from storage is not correct")
}
