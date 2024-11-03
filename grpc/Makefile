# Compile protobuf files and generate Go, gRPC, Gateway, and Swagger files
proto:
	rm -rf pb/customer/*.go
	rm -rf pb/account/*.go
	rm -rf doc/swagger/*.swagger.json

	# Compile customer .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=core_banking \
        proto/manage/customer/*.proto

	# Compile account .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=core_banking \
        proto/manage/account/*.proto

	# Compile manage .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=core_banking \
        proto/manage/*.proto

	# Compile auth .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=core_banking \
        proto/auth/*.proto

	# Compile transactions .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=core_banking \
        proto/transactions/*.proto
	
# PHONY targets (non-file related targets)
.PHONY: proto