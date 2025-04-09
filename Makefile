generate-user-api:
	mkdir -p generated/go/user/v1
	protoc --proto_path=api/user/v1 \
		--go_out=generated/go/user/v1 --go_opt=paths=source_relative \
		--go-grpc_out=generated/go/user/v1 --go-grpc_opt=paths=source_relative \
		api/user/v1/user.proto 

load-test-register:
	ghz \
	--insecure \
	--proto=./api/user/v1/user.proto \
	--call=user_v1.AuthService.Register \
	-d '{"username": "test", "password": "test", "info": {"first_name": "test", "last_name": "test"}}' \
	--total=500_000 \
	--rps=2000 \
	-c 20 \
	localhost:50051