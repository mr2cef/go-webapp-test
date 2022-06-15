build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o run_main main.go

run: build
	./run_main
