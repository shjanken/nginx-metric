hello:
	echo "hello"

cover:
	go test -coverprofile=assets/coverage.out ./...

cover-html:
	go tool cover -html=assets/coverage.out
