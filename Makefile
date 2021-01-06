hello:
	echo "hello"

cover:
	go test -coverprofile=assets/coverage.out ./...

cover-html:
	go tool cover -html=assets/coverage.out

test-with-bigfile:
	BIGFILE=true go test ./...
	
test:
	go test -v ./...