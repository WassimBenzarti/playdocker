

run:
	go run main.go

test:
	go test ./...

dev:
	docker run -d -it -w /workspace -v ${pwd}:/workspace  golang:alpine