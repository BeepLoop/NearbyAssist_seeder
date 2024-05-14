build:
	@go build -o bin/seeder cmd/main.go

run:
	@go run cmd/main.go

test:
	@grc go test -v ./...

clean:
	@rm -rf bin/*
