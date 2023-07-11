LOCAL_BIN:=$(CURDIR)/bin

MIGRATIONS_DIR=./database/migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql

.PHONY: dbup
dbup:
	docker compose up -d

.PHONY: dbdown
dbdown:
	docker compose down


.PHONY: migrationup
migrationup:
	chmod +x ./migrate.sh
	./migrate.sh

.PHONY: server
server: migrationup
	go run ./cmd/bot/main.go

.PHONY: test
test:
	$([INFO] Running tests)
	go test ./...

.PHONY: cover
cover:
	go test -v $$(go list ./... | grep -v -E '/pkg/(api)') -covermode=count -coverprofile=/tmp/coverage.out
	go tool cover -html=/tmp/coverage.out
