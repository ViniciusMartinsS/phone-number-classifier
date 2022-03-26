setup:
	@go mod download
	@go get -d gotest.tools/gotestsum
	@go get -d github.com/go-courier/husky/cmd/husky && husky init

docker:
	@docker-compose up --build

run:
	@go run ./main.go

lint:
	@staticcheck ./...

tests:
	@gotestsum -f pkgname ./test/unit...

coverage:
	@go test -v -cover -coverprofile=r.out -coverpkg ./src/... ./test/...
	@go tool cover -html=r.out -o report.html
	@rm -f r.out
	@google-chrome --new-window report.html
