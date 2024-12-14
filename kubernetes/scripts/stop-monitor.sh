#!/bin/bash

# Delete the namespace if it exists
kubectl delete namespace monitoring || echo "Namespace monitoring not found. Skipping..."

# Uninstall the Helm release if it exists
helm -n monitoring uninstall prometheus-grafana-stack || echo "Release prometheus-grafana-stack not found. Skipping..."
