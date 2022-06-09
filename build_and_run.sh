GOARCH=wasm GOOS=js go build -o web/app.wasm 
echo "wasm done"
go build -o run_main main.go 
echo "main done..."
echo "launching main..."
echo "goto http://127.0.0.1:8000"
./run_main