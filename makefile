.DEFAULT_GOAL = build

build:
	@go mod tidy
	@go build -o bin/mdpages ./cmd/mdpages/main.go

run: build
	@MDPAGES_PORT=5000 ./bin/mdpages
