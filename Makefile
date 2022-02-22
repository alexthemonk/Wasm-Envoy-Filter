build:
	cd test-wasm-plugin && tinygo build -o envoy-filter.wasm -target wasm main.go

default:
	build