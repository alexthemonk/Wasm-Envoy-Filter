build:
	cd test-wasm-plugin && tinygo build -o envoy-filter.wasm -scheduler=none -target=wasi main.go
	git add .
	git commit -m "update wasm filter (automated)"
	git push origin master

default:
	build