FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/program

FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/program /go/bin/program
# Run the hello binary.
ENTRYPOINT ["/go/bin/program"]