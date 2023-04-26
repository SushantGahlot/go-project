generate: clean_proto generate_proto

clean_proto:
	rm -rf /api/generated/*

generate_proto:
	protoc \
	--proto_path ./api/proto \
	--go_out ./api/generated \
	--go_opt=paths=source_relative \
	--go-grpc_out ./api/generated \
	--go-grpc_opt paths=source_relative \
	$(shell find ./api/proto -name '*.proto')
