build:
	cd cmd/api/ && go build -o build/go-api-service

run:
	cmd/api/build/go-api-service

test:
	go test -cover -race `go list ./... | grep -v /vendor`

install:
	dep init

update:
	dep ensure