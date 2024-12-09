#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status.

# Start Databases
echo "Starting postgres01..."
minikube ssh "sudo mkdir -p /data/postgres01 && sudo chown -R 70:70 /data/postgres01"
kubectl apply -f infra/infra-stack/postgres01.deploy.yaml
kubectl apply -f infra/infra-stack/postgres01.service.yaml

echo "Starting postgres02..."
minikube ssh "sudo mkdir -p /data/postgres02 && sudo chown -R 70:70 /data/postgres02"
kubectl apply -f infra/infra-stack/postgres02.deploy.yaml
kubectl apply -f infra/infra-stack/postgres02.service.yaml

echo "Starting redis..."
kubectl apply -f infra/infra-stack/redis.deploy.yaml
kubectl apply -f infra/infra-stack/redis.service.yaml

# Start API Services
echo "Starting manage-service..."
kubectl apply -f infra/infra-stack/manage-service.deploy.yaml
kubectl apply -f infra/infra-stack/manage-service.service.yaml

echo "Starting auth-service..."
kubectl apply -f infra/infra-stack/auth-service.deploy.yaml
kubectl apply -f infra/infra-stack/auth-service.service.yaml

echo "Starting transaction-service..."
kubectl apply -f infra/infra-stack/transaction-service.deploy.yaml
kubectl apply -f infra/infra-stack/transaction-service.service.yaml

echo "Starting notification-service..."
kubectl apply -f infra/infra-stack/notification-service.deploy.yaml

echo "Starting ingress..."
kubectl apply -f infra/infra-stack/ingress.yaml

echo "Infrastructure setup completed!"
