GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOTEST = $(GOCMD) test
GOCLEAN = $(GOCMD) clean
BINARY_FILE = instacrawler

all: run

build:
	$(GOBUILD) -o $(BINARY_FILE) -v .

build_linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_FILE) -v .

build_raspberry:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(BINARY_FILE) -v .

test:
	$(GOTEST) -v .

run:
	@$(GORUN) -v .

clean:
	@$(GOCLEAN)
