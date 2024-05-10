run:
	go run cmd/api/main.go

build all:
	set GOOS=linux
	go build -o bin/Hamzie-API cmd/api/main.go
	set GOOS=windows
	go build -o bin/Hamzie-API.exe cmd/api/main.go

build linux:
	set GOOS=linux
	go build -o bin/Hamzie-API cmd/api/main.go

build win:
	set GOOS=windows
	go build -o bin/Hamzie-API.exe cmd/api/main.go