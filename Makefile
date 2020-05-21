# Go parameters

GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_FILE = instacrawler

all: run

install:
	$(GOGET) -u github.com/joshua-dev/instacrawler

build:
	$(GOBUILD) -o $(BINARY_FILE) -v ./main.go

build_linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_FILE)

test:
	$(GOTEST) -v ./src/...

run:
	@$(GORUN) -v ./src/main.go

clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_FILE)
