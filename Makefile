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
	rm -rf service-auth/config.env
	rm -rf service-management/config.env
	rm -rf service-transfermoney/config.env
	rm -rf service-notification/config.env
	cp config.env service-auth
	cp config.env service-management
	cp config.env service-transfermoney
	cp config.env service-notification

network:
	docker network create bank-network

volume:
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
	docker volume rm auth-database-volume
	docker volume rm core-database-01-volume
	docker volume rm core-database-02-volume
	docker volume rm redis-volume

.PHONY: config network volume start stop rm build build-auth build-management build-transfermoney build-notification

### See network config of any machine
# docker exec -it auth-service ss -tulnp
# docker exec -it auth-service netstat -tulnp