.PHONY: dev run build clean staticcheck fmt pre

dev:
	go run github.com/cespare/reflex -s -r '\.go$$' make clean pre run

run: build
	./particles-demo

build: pre
	go build -o particles-demo cmd/particles-demo/main.go

static/particles.wasm: static/wasm_exec.js
	GOARCH=wasm GOOS=js go build -o static/particles.wasm wasm/main.go

static/wasm_exec.js:
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js static/wasm_exec.js

staticcheck:
	go run honnef.co/go/tools/cmd/staticcheck ./particles/... ./cmd/...

fmt:
	go mod tidy
	go fmt ./particles/... ./cmd/...
	# go vet ./particles/... ./cmd/...

pre: fmt static/particles.wasm

clean:
	rm -rf static/particles.wasm static/wasm_exec.js
