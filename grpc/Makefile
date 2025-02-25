proto:
	rm -rf pb/auth/*.go
	rm -rf pb/cusman/*.go
	rm -rf pb/cusman/customer/*.go
	rm -rf pb/cusman/account/*.go
	rm -rf pb/moneytransfer/*.go
	rm -rf pb/shardman/*.go

	rm -rf doc/auth/swagger/*.swagger.json
	rm -rf doc/cusman/swagger/*.swagger.json
	rm -rf doc/moneytransfer/swagger/*.swagger.json
	rm -rf doc/shardman/swagger/*.swagger.json

	# Compile customer .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/cusman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=cusman \
	proto/cusman/customer/*.proto

	# Compile account .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/cusman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=cusman \
	proto/cusman/account/*.proto

	# Compile cusman .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/cusman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=cusman \
	proto/cusman/*.proto

	# Compile auth .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/auth/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=auth \
	proto/auth/*.proto

	# Compile moneytransfer .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/moneytransfer/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=moneytransfer \
	proto/moneytransfer/*.proto

	# Compile shardman .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/shardman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=shardman \
	proto/shardman/*.proto

auth:
	rm -rf pb/auth/*.go
	rm -rf doc/auth/swagger/*.swagger.json

	# Compile auth .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/auth/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=auth \
	proto/auth/*.proto

cusman:
	rm -rf pb/cusman/*.go
	rm -rf pb/cusman/customer/*.go
	rm -rf pb/cusman/account/*.go
	rm -rf doc/cusman/swagger/*.swagger.json

	# Compile customer .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/cusman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=cusman \
	proto/cusman/customer/*.proto

	# Compile account .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/cusman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=cusman \
	proto/cusman/account/*.proto

	# Compile cusman .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/cusman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=cusman \
	proto/cusman/*.proto

moneytransfer:
	rm -rf pb/moneytransfer/*.go
	rm -rf doc/moneytransfer/swagger/*.swagger.json

	# Compile moneytransfer .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/moneytransfer/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=moneytransfer \
	proto/moneytransfer/*.proto

shardman:
	rm -rf pb/shardman/*.go
	rm -rf doc/shardman/swagger/*.swagger.json

	# Compile shardman .proto files
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/shardman/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=shardman \
	proto/shardman/*.proto

swagger-ui:
	cp -r doc/swagger doc/auth
	cp -r doc/swagger doc/cusman
	cp -r doc/swagger doc/moneytransfer
	cp -r doc/swagger doc/shardman

statik:
	rm -rf ./doc/auth/statik
	rm -rf ./doc/cusman/statik
	rm -rf ./doc/moneytransfer/statik
	rm -rf ./doc/shardman/statik

	statik -src=./doc/auth/swagger -dest=./doc/auth -f
	statik -src=./doc/cusman/swagger -dest=./doc/cusman -f
	statik -src=./doc/moneytransfer/swagger -dest=./doc/moneytransfer -f
	statik -src=./doc/shardman/swagger -dest=./doc/shardman -f
	

.PHONY: proto auth cusman moneytransfer shardman swagger-ui statik