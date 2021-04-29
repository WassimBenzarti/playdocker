

run:
  go run main.go

dev:
  docker run -d -it -w /workspace -v ${pwd}:/workspace  golang:alpine