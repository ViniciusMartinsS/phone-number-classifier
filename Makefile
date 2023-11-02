setup:
	@go mod download

docker:
	@docker-compose up --build

run:
	@go run ./main.go

lint:
	@staticcheck ./...

tests:
	@go test -v ./...

coverage:
	@go test -v -cover -coverprofile=r.out -coverpkg ./src/... ./test/...
	@go tool cover -html=r.out -o report.html
	@rm -f r.out
	@google-chrome --new-window report.html
