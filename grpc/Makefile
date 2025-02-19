proto:
	rm -rf pb/auth/*.go
	rm -rf pb/management/*.go
	rm -rf pb/management/customer/*.go
	rm -rf pb/management/account/*.go
	rm -rf pb/transfermoney/*.go
	rm -rf doc/swagger/*.swagger.json

	# Compile customer .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=management \
	proto/management/customer/*.proto

	# Compile account .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=management \
	proto/management/account/*.proto

	# Compile management .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=management \
	proto/management/*.proto

	# Compile auth .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=auth \
	proto/auth/*.proto

	# Compile transfermoney .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=transfermoney \
	proto/transfermoney/*.proto

.PHONY: proto