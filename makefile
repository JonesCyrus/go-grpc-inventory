proto: ## generate the protobuf for the project
	protoc pkg/pb/inventory.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative

start: ## run the server
	go run cmd/main.go