
start-infra:
	docker compose -f docker-compose.yml up -d

stop-infra:
	docker compose -f docker-compose.yml down

remove-services:
	docker stop manage-service && \
	docker rm manage-service && \
	docker rmi docker-manage-service && \
	docker stop transaction-service && \
	docker rm transaction-service && \
	docker rmi docker-transaction-service && \
	docker stop notification-service && \
	docker rm notification-service && \
	docker rmi docker-notification-service

remove-manage:
	docker stop manage-service && \
	docker rm manage-service && \
	docker rmi docker-manage-service

remove-auth:
	docker stop auth-service && \
	docker rm auth-service && \
	docker rmi docker-auth-service
	
remove-transaction:
	docker stop transaction-service && \
	docker rm transaction-service && \
	docker rmi docker-transaction-service

remove-notification:
	docker stop notification-service && \
	docker rm notification-service && \
	docker rmi docker-notification-service

.PHONY: start-infra stop-infra remove-manage remove-auth remove-transaction