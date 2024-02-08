ifeq ($(OS),Windows_NT)
	AIR_CONFIG := .air-windows.toml
	EXECUTABLE := main.exe
else
	AIR_CONFIG := .air-linux.toml
	EXECUTABLE := main
endif

generate:
	@templ generate
	@pnpm build

build: generate
	@go build -o bin/$(EXECUTABLE) cmd/site/main.go

watch-css:
	@pnpm watch

watch-go:
	@air -c $(AIR_CONFIG)

dev: build
	@make -j 2 watch-css watch-go