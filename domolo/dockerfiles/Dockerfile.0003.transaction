# Build stage
FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN cd services/transaction-service && go mod download
RUN cd services/transaction-service && go build -o main main.go

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/services/transaction-service/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .

# Ensure wait-for.sh has execute permissions
RUN chmod +x ./wait-for.sh ./start.sh

EXPOSE 8081 8091
CMD [ "services/transaction-service/main" ]
ENTRYPOINT [ "/app/start.sh" ]