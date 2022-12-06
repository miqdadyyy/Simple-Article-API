REPO = github.com/miqdadyyy
DB_DSN ?= "mysql://miqdad:anone@tcp(127.0.0.1:3306)/devjets_assessment"

http:
	GOPRIVATE=${REPO} go run ./cmd/http/

migrate:
	migrate -database '${DB_DSN}' -path ./tools/migrations up

wire-gen:
	wire gen ./...

lint:
	golangci-lint run ./...

coverage-clear:
	rm -rf coverage.out coverage.log

coverage-total:
	go tool cover -func coverage.out | grep total | awk '{print "\nTotal coverage: ", $$3}'

coverage-test:
	go test -race `go list ./... | grep -v "cmd/" | grep -v "pkg/" | grep -v "tools/" | grep -v "internal/constant" | grep -v "internal/entity"` -coverprofile=coverage.out

cover: coverage-clear coverage-test coverage-total

mod:
	go mod download

start: mod migrate http