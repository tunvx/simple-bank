#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status
set -u  # Treat unset variables as an error and exit immediately

echo "Removing all deployments..."
kubectl delete deployments --all

echo "Removing all services..."
kubectl delete services --all

echo "Removing all pods..."
kubectl delete pods --all

echo "Removing all persistent volume claims (PVCs)..."
kubectl delete pvc --all

echo "Removing all persistent volumes (PVs)..."
kubectl delete pv --all

echo "Removing all Horizontal Pod Autoscalers (HPAs)..."
kubectl delete hpa --all

echo "Removing ingress...."
kubectl get ingress
kubectl delete ingress --all

echo "Infrastructure cleanup completed."