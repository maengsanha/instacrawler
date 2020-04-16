# Go parameters

GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_FILE = instacrawler
BINARY_UNIX = $(BINARY_FILE)_unix

all: run clean

install:
	$(GOGET)

build:
	$(GOBUILD) -o $(BINARY_FILE) -v ./src/main.go

test:
	$(GOTEST) -v ./src/...

run:
	@$(GORUN) -v ./src/main.go

clean:
	@$(GOCLEAN)
	@rm -f $(BINARY_FILE)
	@rm -f $(BINARY_UNIX)
