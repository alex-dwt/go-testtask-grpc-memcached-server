package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"

	"github.com/alex-dwt/go-testtask-grpc-memcached-server/internal/service"
)

func TestService_GetData(t *testing.T) {
	tests := []struct {
		name               string
		value              string
		isFound            bool
		isError            bool
		storageGetMockFunc func(storage *service.MockStorage)
	}{
		{
			name:    "OK",
			value:   "some-value",
			isFound: true,
			isError: false,
			storageGetMockFunc: func(storage *service.MockStorage) {
				storage.On("Get", mock.Anything, mock.Anything).Return("some-value", nil).Once()
			},
		},
		{
			name:    "Value not found",
			value:   "",
			isFound: false,
			isError: false,
			storageGetMockFunc: func(storage *service.MockStorage) {
				storage.On("Get", mock.Anything, mock.Anything).Return("", service.ErrKeyNotFound).Once()
			},
		},
		{
			name:    "StorageGet returns error",
			value:   "",
			isFound: false,
			isError: true,
			storageGetMockFunc: func(storage *service.MockStorage) {
				storage.On("Get", mock.Anything, mock.Anything).Return("", errors.New("err")).Once()
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storage := &service.MockStorage{}
			tt.storageGetMockFunc(storage)

			svc := service.New(zap.NewNop(), storage)
			value, found, err := svc.GetData(context.Background(), "some-key")

			storage.AssertExpectations(t)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.isFound, found)
				assert.Equal(t, tt.value, value)
			}
		})
	}
}

func TestService_DeleteData(t *testing.T) {
	tests := []struct {
		name                  string
		value                 string
		isFound               bool
		isError               bool
		storageDeleteMockFunc func(storage *service.MockStorage)
	}{
		{
			name:    "OK",
			isError: false,
			storageDeleteMockFunc: func(storage *service.MockStorage) {
				storage.On("Delete", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name:    "StorageDelete returns error",
			isError: true,
			storageDeleteMockFunc: func(storage *service.MockStorage) {
				storage.On("Delete", mock.Anything, mock.Anything).Return(errors.New("err")).Once()
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storage := &service.MockStorage{}
			tt.storageDeleteMockFunc(storage)

			svc := service.New(zap.NewNop(), storage)
			err := svc.DeleteData(context.Background(), "some-key")

			storage.AssertExpectations(t)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_SaveData(t *testing.T) {
	tests := []struct {
		name               string
		value              string
		isFound            bool
		isError            bool
		storageSetMockFunc func(storage *service.MockStorage)
	}{
		{
			name:    "OK",
			isError: false,
			storageSetMockFunc: func(storage *service.MockStorage) {
				storage.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name:    "StorageSet returns error",
			isError: true,
			storageSetMockFunc: func(storage *service.MockStorage) {
				storage.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("err")).Once()
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storage := &service.MockStorage{}
			tt.storageSetMockFunc(storage)

			svc := service.New(zap.NewNop(), storage)
			err := svc.SaveData(context.Background(), "some-key", "some-data")

			storage.AssertExpectations(t)

			if tt.isError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
