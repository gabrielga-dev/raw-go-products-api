#!/bin/bash
set -e

echo "Aplicando ConfigMap e Secret..."
kubectl apply -f 01-configmap-backend.yml
kubectl apply -f 02-secret-backend.yml

echo "Aplicando Postgres..."
kubectl apply -f 03-deployment-postgres.yml
kubectl apply -f 04-service-postgres.yml

echo "Aplicando backend..."
kubectl apply -f 05-deployment-backend.yml
kubectl apply -f 06-service-backend.yml

echo "Recursos aplicados. Verificando pods:"
kubectl get pods -o wide
