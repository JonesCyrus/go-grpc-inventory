proto: ## generate the protobuf for the project
	protoc pkg/pb/api.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

server: ## run the server
	go run cmd/main.go