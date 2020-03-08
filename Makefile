# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_FILE = executeme
BINARY_UNIX = $(BINARY_FILE)_unix

all: run clean

build:
	$(GOBUILD) -o $(BINARY_FILE) -v

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FILE)
	rm -f $(BINARY_UNIX)

run:
	$(GOBUILD) -o $(BINARY_FILE) -v ./...
	./$(BINARY_FILE)


# Cross compilation

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

docker-build:
  docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_FILE)" -v
