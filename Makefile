.PHONY: migrate
migrate:
	go run ./cmd/migrate/main.go

.PHONY: server
server:
	go run ./cmd/server/main.go

.PHONY: test
test:
	go test ./...