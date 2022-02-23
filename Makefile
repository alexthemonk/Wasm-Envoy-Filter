build:
	cd test-wasm-plugin && tinygo build -o envoy-filter.wasm -scheduler=none -target=wasi main.go

default:
	build