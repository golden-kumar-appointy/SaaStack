.PHONY: build

build:
	@echo "Building Core server"
	@go build -o ./bin/core core/main.go

	@echo "Building Core HTTP Proxy server"
	@go build -o ./bin/http-proxy http-gateway/main.go

	@echo "Building Custom Payment Plugin server"
	@go build -o ./bin/plugins/payment plugins/payment/custom/main.go

	@echo "Building Custom Email Plugin server"
	@go build -o ./bin/plugins/email plugins/email/custom/main.go

