GOCMD = go
GOBUILD = $(GOCMD) build
GORUN = $(GOCMD) run
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
BINARY_FILE = instacrawler
RM = rm -f

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
	@$(RM) $(BINARY_FILE)
