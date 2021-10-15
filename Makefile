.PHONY:
.SILENT:

test:
	env GO111MODULE=on go test --short -coverprofile=cover.out -v ./...
	make test.coverage

test.coverage:
	env GO111MODULE=on go tool cover -func=cover.out

lint:
	golangci-lint run
