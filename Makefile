BINARY=ssh-portfolio
CMD=./cmd/server

.PHONY: build test lint vet run clean keygen

build:
	go build -o $(BINARY) $(CMD)

test:
	go test -race -cover ./...

lint:
	golangci-lint run ./...

vet:
	go vet ./...

run: build keygen
	./$(BINARY) -port 2222 -host-key ./data/host_key

keygen:
	@mkdir -p data
	@if [ ! -f data/host_key ]; then \
		ssh-keygen -t ed25519 -f data/host_key -N "" -C "ssh-portfolio" && \
		echo "Host key generated at data/host_key"; \
	else \
		echo "Host key already exists, skipping."; \
	fi

clean:
	rm -f $(BINARY) coverage.out coverage.html

cover:
	go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report at coverage.html"
