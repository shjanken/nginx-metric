GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_NAME=ngxmetric
BINARY_PATH=target
TESTFILE=testdata/access.500.log
RELEASE_PATH=release
TARGET_PATH=target

all: test build
build: clean
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
	rm -rf $(TARGET_PATH) 
	rm -rf $(RELEASE_PATH)
release: build
	mkdir $(RELEASE_PATH)
	tar -czf $(RELEASE_PATH)/ngxmetric.tar.gz $(TARGET_PATH)/*
run: build
	./$(BINARY_PATH)/$(BINARY_NAME) $(TESTFILE)
# docker-compse
init-db:
	docker-compose up -d
