tag=0.0.3

run:
	go run main.go

test:
	go test ./...

build:
	bash scripts/build.sh playdocker


release:
	gh release create ${tag} build/* --title "${tag}"

dev:
	docker run -d -it -w /workspace -v ${pwd}:/workspace  golang:alpine