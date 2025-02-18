#!/bin/bash

# Step 1: Get the ClusterIP of the Ingress NGINX Controller
INGRESS_IP=$(kubectl get svc -n ingress-nginx ingress-nginx-controller -o jsonpath='{.spec.clusterIP}')

# Check if we got the IP
if [ -z "$INGRESS_IP" ]; then
  echo "Ingress controller IP not found. Please check if the Ingress service is available."
  exit 1
fi

echo "Ingress controller ClusterIP: $INGRESS_IP"

# Step 2: Remove old entries from /etc/hosts
echo "Removing old entries from /etc/hosts..."
sudo sed -i '' '/auth-service.banking.com/d' /etc/hosts
sudo sed -i '' '/manage-service.banking.com/d' /etc/hosts
sudo sed -i '' '/transaction-service.banking.com/d' /etc/hosts
sudo sed -i '' '/auth-grpc-service.banking.com/d' /etc/hosts
sudo sed -i '' '/manage-grpc-service.banking.com/d' /etc/hosts
sudo sed -i '' '/transaction-grpc-service.banking.com/d' /etc/hosts

# Step 3: Update /etc/hosts with new entries
echo "Updating /etc/hosts..."

# Add the entries for auth-service, manage-service, and transaction-service
echo "$INGRESS_IP auth-service.banking.com" | sudo tee -a /etc/hosts
echo "$INGRESS_IP manage-service.banking.com" | sudo tee -a /etc/hosts
echo "$INGRESS_IP transaction-service.banking.com" | sudo tee -a /etc/hosts

echo "$INGRESS_IP auth-grpc-service.banking.com" | sudo tee -a /etc/hosts
echo "$INGRESS_IP manage-grpc-service.banking.com" | sudo tee -a /etc/hosts
echo "$INGRESS_IP transaction-grpc-service.banking.com" | sudo tee -a /etc/hosts

echo "/etc/hosts updated successfully."
