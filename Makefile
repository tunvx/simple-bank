build-auth:
	docker build -t auth-service -f service-auth/Dockerfile .

build-management:
	docker build -t management-service -f service-management/Dockerfile .

build-transfermoney:
	docker build -t transfermoney-service -f service-transfermoney/Dockerfile .

build-notification:
	docker build -t notification-service -f service-notification/Dockerfile .

build:
	docker build -t auth-service -f service-auth/Dockerfile .
	docker build -t management-service -f service-management/Dockerfile .
	docker build -t transfermoney-service -f service-transfermoney/Dockerfile .
	docker build -t notification-service -f service-notification/Dockerfile .

rm-build:
	docker image rm auth-service
	docker image rm management-service
	docker image rm transfermoney-service
	docker image rm notification-service

config:
	cp config.dev.env service-lookup/config.env
	cp config.dev.env service-auth/config.env
	cp config.dev.env service-management/config.env
	cp config.dev.env service-transfermoney/config.env
	cp config.dev.env service-notification/config.env

network:
	docker network create bank-network

volume:
	docker volume create original-database-volume
	docker volume create auth-database-volume
	docker volume create core-database-01-volume
	docker volume create core-database-02-volume
	docker volume create redis-volume

start:
	docker compose --env-file .env -f docker-compose.yml up -d 

stop:
	docker compose --env-file .env -f docker-compose.yml down

rm: 
	docker network rm bank-network
	docker volume rm original-database-volume
	docker volume rm auth-database-volume
	docker volume rm core-database-01-volume
	docker volume rm core-database-02-volume
	docker volume rm redis-volume

.PHONY: config network volume start stop rm build build-auth build-management build-transfermoney build-notification

### See network config of any machine
# docker exec -it auth-service ss -tulnp
# docker exec -it auth-service netstat -tulnp