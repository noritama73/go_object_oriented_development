GO_VERSION=$(cat .go-version)

vendor:
	@go mod vendor

fmt:
	@go fmt ./...

build: vendor
	@go build -o main main.go

test: vendor
	@go test ./... -v -cover -coverprofile=cover.out

coverage:
	@go tool cover -html=cover.out -o cover.html
