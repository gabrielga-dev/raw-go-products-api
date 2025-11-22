# Raw Go Products API

This project uses raw go to build a web application to create, update, list and delete products.

## Tecnologies
- Go (v1.22.2)
- Postgres (v16)
- Kubernetes

## How to run it?
You can run the project with and without kubernetes.

### Without Kubernetes
1. Run the `compose.yml` file with:
> docker compose up -d

So the database will be up.

2. Set up env vars

| Env var name | Value 
|----------|----------|
| DATABASE_HOST | localhost |
| DATABASE_PORT | 5432 |
| DATABASE_NAME | go-products |
| DATABASE_USER | admin |
| DATABASE_PASSWORD | admin123 |

3. Run go project with:
> go run main.go

### With Kubernetes
1. Go to kubernetes dir
> cd ./k8s

2. Run the deploy script
>./deploy.sh

3. Do the port forward command to access the project
> kubectl port-forward deploy/backend-deployment 8000:8000

You can use k9s to make de port forward more easily!

4. To end the project, run the reset script:
>./reset.sh


## Access the project
Once the project is up, access: http://localhost:8000