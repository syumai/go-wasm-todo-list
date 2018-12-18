.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o ./index.wasm .
