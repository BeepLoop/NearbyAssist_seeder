build:
	@go build -o bin/seeder cmd/main.go

build-win:
	@GOOS=windows GOARCH=amd64 go build -o bin/seeder.exe cmd/main.go

run:
	@go run cmd/main.go

test:
	@grc go test -v ./...

clean:
	@rm -rf bin/*
