#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

####################################################
### FOR ORBSTACK Env: orb run bash -c

echo "Creating persistent storage directories and setting permissions..."

# Create the directory for postgres01, set ownership, and secure it with chmod 700
orb run bash -c "sudo mkdir -p /data/postgres01 && sudo chown -R 70:70 /data/postgres01 && sudo chmod 700 /data/postgres01"
echo "Persistent storage for postgres01 has been created, ownership set, and secured."

# Create the directory for postgres02, set ownership, and secure it with chmod 700
orb run bash -c "sudo mkdir -p /data/postgres02 && sudo chown -R 70:70 /data/postgres02 && sudo chmod 700 /data/postgres02"
echo "Persistent storage for postgres02 has been created, ownership set, and secured."

echo "Persistent data setup completed successfully!"


####################################################
### FOR MINIKUBE Env: minikube ssh

# echo "Creating persistent storage directories and setting permissions..."

# # Create the directory for postgres01, set ownership, and secure it with chmod 700
# minikube ssh "sudo mkdir -p /data/postgres01 && sudo chown -R 70:70 /data/postgres01 && sudo chmod 700 /data/postgres01"
# echo "Persistent storage for postgres01 has been created, ownership set, and secured."

# # Create the directory for postgres02, set ownership, and secure it with chmod 700
# minikube ssh "sudo mkdir -p /data/postgres02 && sudo chown -R 70:70 /data/postgres02 && sudo chmod 700 /data/postgres02"
# echo "Persistent storage for postgres02 has been created, ownership set, and secured."

# echo "Persistent data setup completed successfully!"
