APP_NAME := passpop
VERSION  := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")

build:
	go build -o $(APP_NAME) main.go
	@echo "✅ Built binary: ./$(APP_NAME)"

test:
	go test ./... -v

clean:
	rm -f $(APP_NAME)
	rm -rf dist
	@echo "🧹 Cleaned build artifacts."

release:
	goreleaser release --snapshot --clean
	@echo "📦 Local release built in ./dist"

tag:
ifndef v
	$(error Please provide a version with: make tag v=1.2.3)
endif
	git tag v$(v)
	git push origin v$(v)
	@echo "🏷️  Tagged version: v$(v)"

.PHONY: build test clean release tag