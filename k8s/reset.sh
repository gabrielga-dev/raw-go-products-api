#!/bin/bash
set -e
echo "Deletando backend e postgres, configmap e secret..."
kubectl delete -f 05-deployment-backend.yml --ignore-not-found
kubectl delete -f 06-service-backend.yml --ignore-not-found
kubectl delete -f 03-deployment-postgres.yml --ignore-not-found
kubectl delete -f 04-service-postgres.yml --ignore-not-found
kubectl delete -f 01-configmap-backend.yml --ignore-not-found
kubectl delete -f 02-secret-backend.yml --ignore-not-found

echo "Feito. Recursos removidos."
