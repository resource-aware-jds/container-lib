generate:
	protoc --go_out=./generated/proto/ \
        		--go-grpc_out=./generated/proto/ \
        		./proto/* & \
	go generate ./...
