godev:
	GOPROXY=off go run main.go

check:
	GOPROXY=off go fmt ./...
	GOPROXY=off go test ./... -cover

coverprofile:
	GOPROXY=off go test ./... -coverprofile=c.out && GOPROXY=off go tool cover -html=c.out

linux:
	GOPROXY=off GOOS=linux GOARCH=amd64 go build -o gin-demo main.go

mac:
	GOPROXY=off go build -o gin-demo main.go