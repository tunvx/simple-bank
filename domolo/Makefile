
start-infra:
	docker compose -f docker-compose.yml up -d

stop-infra:
	docker compose -f docker-compose.yml down

start-monitor:
	docker compose -f docker-compose.monitoring.yml up -d

stop-monitor:
	docker compose -f docker-compose.monitoring.yml down

rm-transaction:
	docker stop transaction-service && \
	docker rm transaction-service && \
	docker rmi domolo-transaction-service

.PHONY: start-infra stop-infra start-monitor stop-monitor