generate-user-api:
	mkdir -p generated/go/user/v1
	protoc --proto_path=api/user/v1 \
		--go_out=generated/go/user/v1 --go_opt=paths=source_relative \
		--go-grpc_out=generated/go/user/v1 --go-grpc_opt=paths=source_relative \
		api/user/v1/user.proto 
