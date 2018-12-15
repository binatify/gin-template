test:
	GOPROXY=off go test ./... -cover
format:
	GOPROXY=off go fmt ./...
install:
	go fmt ./...
lint:
	golint ./... | grep -v ^vendor/ | grep -v 'exported'
cover:
	GOPROXY=off go test ./... -coverprofile=c.out && GOPROXY=off go tool cover -html=c.out