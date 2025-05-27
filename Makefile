.PHONY: build

build:
	@go build -o ./bin/core core/main.go
	@go build -o ./bin/http-proxy http-gateway/main.go
	@go build -o ./bin/plugins/payment plugins/payment/custom/main.go
	@go build -o ./bin/plugins/email plugins/email/custom/main.go

