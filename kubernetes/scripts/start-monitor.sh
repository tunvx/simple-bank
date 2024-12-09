#!/bin/bash

# Create the namespace if it doesn't exist
kubectl create namespace monitoring 2>/dev/null || true

# Install Prometheus and Grafana
helm -n monitoring install prometheus-grafana-stack -f monitor/values-prometheus.yaml monitor/kube-prometheus-stack
