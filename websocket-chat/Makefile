build:
	@go build -o bin/ws_chat

run: build
	@./bin/ws_chat

test:
	@go clean -testcache
	@go test -race -v ./...

clean:
	@rm -rf bin


