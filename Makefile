GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_NAME=ngxmetric
BINARY_PATH=target
TESTFILE=testdata/access.500.log

all: test build
build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) cmd/ngxmetric/main.go
	cp -r config $(BINARY_PATH)/
cover:
	$(GOTEST) -coverprofile=assets/coverage.out ./...
cover-html: cover
	go tool cover -html=assets/coverage.out
test-with-bigfile:
	BIGFILE=true $(GOTEST) ./...
test:
	go test -v ./...
clean:
	$(GOCLEAN)
	rm -rf target/$(BINARY_NAME)
run: build
	./$(BINARY_PATH)/$(BINARY_NAME) $(TESTFILE)
# docker-compse
init-db:
	docker-compose up -d
