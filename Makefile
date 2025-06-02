.PHONY: build generate clean

all: clean generate build

build:
	@echo "Building from main.go after reading config.yaml"
	@go build -o ./bin/ main.go

generate:
	@echo "Generating from email proto files"
	@cd interfaces/email && buf generate && cd ../../
	@echo "Generating from payment proto files"
	@cd interfaces/payment && buf generate  && cd ../../

clean:
	@echo "Cleaning proto generate files"
	@rm -r gen/
	@echo "Cleaning generate docs"
	@rm -r docs/email/
	@rm -r docs/payment/
