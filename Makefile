.PHONY: proto
proto:
	protoc \
	-I=./proto \
	--go_out=./proto \
	--go_opt=paths=source_relative \
	./proto/sample.proto

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: test
test:
	go test -race -v ./...

.PHONY: dev
dev:
	go run ./...
