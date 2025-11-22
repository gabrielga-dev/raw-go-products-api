# syntax=docker/dockerfile:1

########## STAGE 1 ##########
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o raw-go-products-api .

########## STAGE 2 ##########
FROM gcr.io/distroless/static

WORKDIR /

COPY --from=builder /app/raw-go-products-api .

COPY --from=builder /app/templates ./templates
COPY --from=builder /app/migrations ./migrations

EXPOSE 8000

ENTRYPOINT ["/raw-go-products-api"]
