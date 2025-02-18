.PHONY: vm-persist-data, rm-persist-data, start-infra, update-etc-hosts, stop-infra, start-monitor, stop-monitor, test-connection \
postgres01, postgres02, redis, manage-service, auth-service, notification-service, ingress \
rm-postgres01, rm-postgres02, rm-redis, rm-manage-service, rm-auth-service, rm-transaction-service, rm-notification-service, rm-ingress

vm-persist-data:
	./scripts/vm-persist-data.sh

rm-persist-data:
	orb run bash -c "sudo rm -rf /data/*"

start-infra: FORCE
	./scripts/start-infra.sh
FORCE:

update-etc-hosts:
	./scripts/update-etc-hosts.sh

stop-infra:	
	./scripts/stop-infra.sh

start-monitor: FORCE
	./scripts/start-monitor.sh
FORCE:

stop-monitor:
	./scripts/stop-monitor.sh

test-connection:
	curl -X GET http://auth-service.banking.com/v1/generate_test_access_token



######################################################################################################
######################################################################################################
### FOR ORBSTACK Env: orb run bash -c
### FOR MINIKUBE Env: minikube ssh

postgres01:
# 	Creates the required directory and permissions for postgres01, then deploys its Kubernetes resources.
# 	minikube ssh "sudo mkdir -p /data/postgres01 && sudo chown -R 70:70 /data/postgres01 && sudo chmod 700 /data/postgres01"
	orb run bash -c "sudo mkdir -p /data/postgres01 && sudo chown -R 70:70 /data/postgres01 && sudo chmod 700 /data/postgres01"
	kubectl apply -f infra/infra-stack/postgres01.deploy.yaml
	kubectl apply -f infra/infra-stack/postgres01.service.yaml

rm-postgres01:
#   Deletes deployment, service, PVC, and PV for postgres01.
	kubectl delete deployment postgres01 --ignore-not-found
	kubectl delete svc postgres01 --ignore-not-found
	kubectl delete pvc postgres01-pvc --ignore-not-found
	kubectl delete pv postgres01-pv --ignore-not-found


postgres02:
#   Creates the required directory and permissions for postgres02, then deploys its Kubernetes resources.
#   minikube ssh "sudo mkdir -p /data/postgres02 && sudo chown -R 70:70 /data/postgres02 && sudo chmod 700 /data/postgres02"
	orb run bash -c "sudo mkdir -p /data/postgres02 && sudo chown -R 70:70 /data/postgres02 && sudo chmod 700 /data/postgres02"
	kubectl apply -f infra/infra-stack/postgres02.deploy.yaml
	kubectl apply -f infra/infra-stack/postgres02.service.yaml

rm-postgres02:
#   Deletes deployment, service, PVC, and PV for postgres02.
	kubectl delete deployment postgres02 --ignore-not-found
	kubectl delete svc postgres02 --ignore-not-found
	kubectl delete pvc postgres02-pvc --ignore-not-found
	kubectl delete pv postgres02-pv --ignore-not-found


redis:
#   Deploys the Redis service and its associated deployment.
	kubectl apply -f infra/infra-stack/redis.deploy.yaml && \
	kubectl apply -f infra/infra-stack/redis.service.yaml

rm-redis:
#   Deletes deployment, service, PVC, and PV for Redis.
	kubectl delete deployment redis --ignore-not-found
	kubectl delete svc redis --ignore-not-found
	kubectl delete pvc redis-pvc --ignore-not-found
	kubectl delete pv redis-pv --ignore-not-found


manage-service:
#   Deploys the manage-service and its associated Kubernetes resources.
	kubectl apply -f infra/infra-stack/manage-service.deploy.yaml && \
	kubectl apply -f infra/infra-stack/manage-service.service.yaml

rm-manage-service:
	# Deletes deployment and service for manage-service.
	kubectl delete deployment manage-service --ignore-not-found
	kubectl delete svc manage-service --ignore-not-found


auth-service:
#   Deploys the auth-service and its associated Kubernetes resources.
	kubectl apply -f infra/infra-stack/auth-service.deploy.yaml && \
	kubectl apply -f infra/infra-stack/auth-service.service.yaml

rm-auth-service:
#   Deletes deployment, service, PVC, and PV for auth-service.
	kubectl delete deployment auth-service --ignore-not-found
	kubectl delete svc auth-service --ignore-not-found
	kubectl delete pvc auth-service-pvc --ignore-not-found
	kubectl delete pv auth-service-pv --ignore-not-found


transaction-service:
#   Deploys the transaction-service and its associated Kubernetes resources.
	kubectl apply -f infra/infra-stack/transaction-service.deploy.yaml && \
	kubectl apply -f infra/infra-stack/transaction-service.service.yaml && \
	kubectl apply -f infra/infra-stack/transaction-service.hpa.yaml

rm-transaction-service:
#   Deletes deployment, service, PVC, PV, and HPA for transaction-service.
	kubectl delete deployment transaction-service --ignore-not-found
	kubectl delete svc transaction-service --ignore-not-found
	kubectl delete pvc transaction-service-pvc --ignore-not-found
	kubectl delete pv transaction-service-pv --ignore-not-found
	kubectl delete hpa transaction-service-hpa --ignore-not-found


notification-service:
#   Deploys the notification-service with its Kubernetes resources.
	kubectl apply -f infra/infra-stack/notification-service.deploy.yaml

rm-notification-service:
#   Deletes deployment, service, PVC, and PV for notification-service.
	kubectl delete deployment notification-service --ignore-not-found
	kubectl delete svc notification-service --ignore-not-found
	kubectl delete pvc notification-service-pvc --ignore-not-found
	kubectl delete pv notification-service-pv --ignore-not-found

ingress:
	kubectl apply -f infra/infra-stack/ingress.yaml

rm-ingress:
	kubectl get ingress
	kubectl delete ingress --all


### START MINIKUBE (FOR LOCAL)
# minikube start --cpus=10 --memory=13600 --addons=metrics-server,ingress

### EXPOSE PORT TO LOCAL (FOR MAC)
# kubectl get svc
# kubectl get ingress
# curl -X GET http://auth-service.banking.com/v1/generate_test_access_token

### MOCK DATA FOR TESTTING
# kubectl port-forward svc/postgres01 5432:5432
# kubectl port-forward svc/postgres02 5433:5432
# make ...

### DESCRIBE POD & NODE
# kubectl logs .pid
# kubectl describe pod  .pid
# kubectl describe nodes