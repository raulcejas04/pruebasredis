#https://github.com/codefresh-contrib/golang-sample-app
FROM golang:1.16-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/go-sample-app .


# This container exposes port 8080 to the outside world
EXPOSE 5000

# Run the binary program produced by `go install`
#CMD ["./out/go-sample-app"]


