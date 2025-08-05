run:
	cd ./cmd && export APP_ENV=dev && go run main.go

mocks:
	mockgen -source=internal/handler/message_handler.go -destination=internal/mocks/message_handler.go -package mocks
	mockgen -source=internal/handler/scheduler_handler.go -destination=internal/mocks/scheduler_handler.go -package mocks
	mockgen -source=internal/service/message_service.go -destination=internal/mocks/message_service.go -package mocks
	mockgen -source=internal/scheduler/scheduler.go -destination=internal/mocks/scheduler.go -package mocks

tests:
	go test -v ./internal/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html