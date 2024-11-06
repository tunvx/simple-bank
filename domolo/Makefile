
start-infra:
	docker compose -f docker-compose.yml up -d

stop-infra:
	docker compose -f docker-compose.yml down

start-monitor:
	docker compose -f docker-compose.monitoring.yml up -d

stop-monitor:
	docker compose -f docker-compose.monitoring.yml down

.PHONY: start-infra stop-infra start-monitor stop-monitor