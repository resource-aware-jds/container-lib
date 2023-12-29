generate:
	protoc --go_out=./generated/proto/ \
        		--go-grpc_out=./generated/proto/ \
        		--proto_path=./proto \
        		./proto/* & \
	go generate ./...
