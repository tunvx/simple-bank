IMAGE_PREFIX=tunvx/
VERSION=v0.0.1

.PHONY: config network volume start-infra stop-infra clear-infra start-services stop-services\ 
		build rm-build build-auth build-management build-transfermoney build-notification

config:
	cp config.dev.env vnb-auth-service/config.env
	cp config.dev.env vnb-cusman-service/config.env
	cp config.dev.env vnb-moneytransfer-service/config.env
	cp config.dev.env vnb-shardman-service/config.env

network:
	docker network create bank-network

volume:
	docker volume create original-database-volume
	docker volume create auth-database-volume
	docker volume create core-database-01-volume
	docker volume create core-database-02-volume
	docker volume create redis-volume
	docker volume create kafka-volume

start-infra:
	docker compose --env-file .env -f docker-compose.infra.yml up -d 

stop-infra:
	docker compose --env-file .env -f docker-compose.infra.yml down

clear-infra:
	docker network rm bank-network
	docker volume rm original-database-volume
	docker volume rm auth-database-volume
	docker volume rm core-database-01-volume
	docker volume rm core-database-02-volume
	docker volume rm redis-volume
	docker volume rm kafka-volume

start-services:
	docker compose --env-file .env -f docker-compose.services.yml up -d 

stop-services:
	docker compose --env-file .env -f docker-compose.services.yml down

# Build image for each service
build-auth:
	docker build -t $(IMAGE_PREFIX)vnb-auth-service:$(VERSION) -f vnb-auth-service/Dockerfile .

build-cusman:
	docker build -t $(IMAGE_PREFIX)vnb-cusman-service:$(VERSION) -f vnb-cusman-service/Dockerfile .

build-moneytransfer:
	docker build -t $(IMAGE_PREFIX)vnb-moneytransfer-service:$(VERSION) -f vnb-moneytransfer-service/Dockerfile .

build-shardman:
	docker build -t $(IMAGE_PREFIX)vnb-shardman-service:$(VERSION) -f vnb-shardman-service/Dockerfile .

# Remove image of each service
rm-build-auth:
	docker image rm $(IMAGE_PREFIX)vnb-auth-service:$(VERSION)

rm-build-cusman:
	docker image rm $(IMAGE_PREFIX)vnb-cusman-service:$(VERSION)

rm-build-moneytransfer:
	docker image rm $(IMAGE_PREFIX)vnb-moneytransfer-service:$(VERSION)

rm-build-shardman:
	docker image rm $(IMAGE_PREFIX)vnb-shardman-service:$(VERSION)

# Build image for all services
build: build-auth build-cusman build-moneytransfer build-shardman

# Remove image of all services
rm-build: rm-build-auth rm-build-cusman rm-build-moneytransfer rm-build-shardman

### See network config of any machine
# docker exec -it auth-service ss -tulnp
# docker exec -it auth-service netstat -tulnp