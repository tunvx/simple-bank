#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status.

# Step 1: Deploy the Ingress NGINX Controller
echo "Starting ingress-nginx..."
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.1/deploy/static/provider/cloud/deploy.yaml
echo "Waiting for 15 seconds to allow Ingress NGINX Controller initialization..."
sleep 15  # Allow time forIngress NGINX Controller to be fully up and running

# Step 2: Deploy Database, Caching, and MQ Services

# Start Postgres01 (core-database)
echo "Starting postgres01..."
kubectl apply -f infra/infra-stack/postgres01.deploy.yaml
kubectl apply -f infra/infra-stack/postgres01.service.yaml

# Start Postgres02 (auth-database)
echo "Starting postgres02..."
kubectl apply -f infra/infra-stack/postgres02.deploy.yaml
kubectl apply -f infra/infra-stack/postgres02.service.yaml

# Start Redis (Caching, MQ services)
echo "Starting redis..."
kubectl apply -f infra/infra-stack/redis.deploy.yaml
kubectl apply -f infra/infra-stack/redis.service.yaml

# Step 3: Wait for DB services to initialize
echo "Waiting for 15 seconds to allow database initialization..."
sleep 15  # Allow time for DB services to be fully up and running

# Step 4: Deploy Simple Banking API Services

# Start Manage Service
echo "Starting manage-service..."
kubectl apply -f infra/infra-stack/manage-service.deploy.yaml
kubectl apply -f infra/infra-stack/manage-service.service.yaml

# Start Auth Service
echo "Starting auth-service..."
kubectl apply -f infra/infra-stack/auth-service.deploy.yaml
kubectl apply -f infra/infra-stack/auth-service.service.yaml

# Start Transaction Service
echo "Starting transaction-service..."
kubectl apply -f infra/infra-stack/transaction-service.deploy.yaml
kubectl apply -f infra/infra-stack/transaction-service.hpa.yaml  # Horizontal Pod Autoscaler for transaction service
kubectl apply -f infra/infra-stack/transaction-service.service.yaml

# Start Notification Service
echo "Starting notification-service..."
kubectl apply -f infra/infra-stack/notification-service.deploy.yaml

# Step 5: Deploy Ingress
echo "Starting ingress..."
kubectl apply -f infra/infra-stack/ingress.yaml

# Final Message
echo "Infrastructure setup completed!"
