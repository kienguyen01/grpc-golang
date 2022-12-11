gen:
	protoc --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

clean:
	rm pb/proto/*.go

run-server:
	go run cmd/server/main.go -port 8080

run-client:
	go run cmd/client/main.go -address 0.0.0.0:8080

test:
	go test -cover -race ./...