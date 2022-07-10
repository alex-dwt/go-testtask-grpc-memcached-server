# go-testtask-grpc-memcached-server

## Что требуется добавить/изменить/улучшить
- порт сервера захардкожен - 8889  
- нету перехвата SIG_INT/SIG_TERM в основной main  
- нету GracefulStop у grpc сервера  
- сейчас данные для сохранения принимаются только как string  
- написаны только юнит тесты  

## Запуск
STORAGE_TYPE=MEMORY go run ./cmd/server/  
STORAGE_TYPE=MEMCACHED go run ./cmd/server/  