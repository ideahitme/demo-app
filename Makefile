default: build

build:
	@go build -o build/demo-app .
clean:
	@rm -rf build/

.PHONY: build clean
