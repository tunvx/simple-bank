rm-all:
	kubectl delete deployments --all && \
	kubectl delete services --all && \
	kubectl delete pods --all && \
	kubectl delete pvc --all && \
	kubectl delete pv --all 

rm-data:
	minikube ssh -- 'sudo rm -rf /data/*'

postgres01:
	kubectl apply -f postgres01.deploy.yaml && \
	kubectl apply -f postgres01.service.yaml

postgres02:
	kubectl apply -f postgres02.deploy.yaml && \
	kubectl apply -f postgres02.service.yaml

redis:
	kubectl apply -f redis.deploy.yaml && \
	kubectl apply -f redis.service.yaml

manage-service:
	kubectl apply -f manage-service.deploy.yaml && \
	kubectl apply -f manage-service.service.yaml

auth-service:
	kubectl apply -f auth-service.deploy.yaml && \
	kubectl apply -f auth-service.service.yaml

transaction-service:
	kubectl apply -f transaction-service.deploy.yaml && \
	kubectl apply -f transaction-service.service.yaml

notification-service:
	kubectl apply -f notification-service.deploy.yaml

.PHONY: rm-all, postgres01, postgres02, redis, manage-service,  auth-service, notification-service

### START MINIKUBE (FOR LOCAL)
# minikube start --cpus=10 --memory=13600 --addons=metrics-server,ingress

### EXPOSE PORT TO LOCAL (FOR MAC)
# kubectl get svc
# minikube addons enable ingress
# minikube tunnel # run "minikube tunnel" and your ingress resources would be available at "127.0.0.1"
# curl -X GET http://auth-service.banking.com/v1/generate_test_access_token

### MOCK DATA FOR TESTTING
# kubectl port-forward svc/postgres01 5432:5432
# kubectl port-forward svc/postgres02 5433:5432
# make ...