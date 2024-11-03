# Makefile

# Variables
DOCKER_COMPOSE_FILE=docker-compose.yml
DATABASE_CONTAINER=banking-server-posgres
DB_USERNAME=root
DB_OWNER=root
DB_NAME=banking_db
DB_PASSWORD=secret
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable
MIGRATION_PATH=./db/migration
MOCKDATA_PATH=./db/mockdata

# Start the docker compose services (infrastructure)
start-infra:
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build -d

# Stop the docker compose services and remove Docker images
stop-infra:
	docker compose -f $(DOCKER_COMPOSE_FILE) down && \
	docker rmi core-banking-services-manager-service && \
	docker rmi core-banking-services-worker-service && \
	docker rmi core-banking-services-customer-service && \
	docker rmi core-banking-services-auth-service && \
	docker rmi core-banking-services-account-service

# stop-infra:
# 	docker compose -f $(DOCKER_COMPOSE_FILE) down && \
# 	docker rmi core-banking-services-manager-service && \
# 	docker rmi core-banking-services-worker-service && \
# 	docker rmi core-banking-services-customer-service && \
# 	docker rmi core-banking-services-auth-service && \
# 	docker rmi core-banking-services-account-service && \
# 	docker volume rm core-banking-services_bank-volume

# Start Redis container
redis:
	docker run --name redis -p 6379:6379 -d redis:7.4-alpine3.20

# Start PostgreSQL container
postgres:
	docker run --name banking-server-posgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.0-alpine3.20

# Create the database
createdb:
	docker exec -it $(DATABASE_CONTAINER) createdb --username=$(DB_USERNAME) --owner=$(DB_OWNER) $(DB_NAME)

# Drop the database
dropdb:
	docker exec -it $(DATABASE_CONTAINER) dropdb $(DB_NAME)

# Run migrations up
migrateup:
	migrate -path $(MIGRATION_PATH) -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -verbose up

# Run migrations down
migratedown:
	migrate -path $(MIGRATION_PATH) -database "postgresql://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" -verbose down

# Generate SQL code using sqlc
gensqlc:
	sqlc generate

# Generate mock data using Python script
mockdata:
	python3 $(MOCKDATA_PATH)/generate_mockdata.py

# Generate Go mocks for tests
mockgen:
	mockgen -package mockdb -destination db/mockgen/store.go github.com/tunvx/banksys/db/sqlc Store

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

	# Compile service .proto file
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=doc/swagger  --openapiv2_opt=allow_merge=true,merge_file_name=core_banking \
        proto/*.proto

# Run the local server
manager-service:
	go run cmd/manager/main.go

worker-service:
	go run cmd/worker/main.go

customer-service:
	go run cmd/customer/main.go

auth-service:
	go run cmd/auth/main.go

account-service:
	go run cmd/account/main.go

# Run Go tests with coverage
testlogic:
	go test -v -cover -short ./...

# Clean Go test cache
cleantestlogic:
	go clean -testcache

# Run K6 load testing scripts
k6loadtesting:
	./load_testing/run_all_test_transfer_money.sh

# PHONY targets (non-file related targets)
.PHONY: start-infra stop-infra redis postgres createdb dropdb migrateup migratedown \
        gensqlc mockdata mockgen proto manager-service worker-service customer-service \
        auth-service account-service testlogic cleantestlogic k6loadtesting