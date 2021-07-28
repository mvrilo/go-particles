.PHONY: dev run build clean staticcheck fmt pre

run: build
	./particles-demo

run-dev: build-dev
	./particles-demo-dev

dev:
	go run github.com/cespare/reflex -s -r '\.go$$' make clean pre run-dev

build-dev: pre
	go build -race -o particles-demo-dev cmd/particles-demo/main.go

build: pre
	go build -o particles-demo cmd/particles-demo/main.go

static/particles.wasm: static/wasm_exec.js
	# can't use tinygo yet, see:
	# https://github.com/tinygo-org/tinygo/issues/1848
	# https://github.com/tinygo-org/tinygo/issues/1979
	# tinygo build -target wasm -o static/particles.wasm wasm/main.go
	GOARCH=wasm GOOS=js go build -o static/particles.wasm wasm/main.go

static/wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js static/wasm_exec.js

staticcheck:
	go run honnef.co/go/tools/cmd/staticcheck ./particles ./cmd/...

fmt:
	@go mod tidy
	@go fmt ./...
	@go vet ./particles ./cmd/...
	@GOARCH=wasm GOOS=js go vet ./canvas ./wasm

pre: static/particles.wasm fmt

clean:
	rm -rf particles-demo particles-demo-dev static/particles.wasm static/wasm_exec.js
