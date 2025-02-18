# Build stage
FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN cd services/manage-service && go mod download
RUN cd services/manage-service && go build -o main main.go

# Run stage
FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/services/manage-service/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY services/manage-service/db/migration ./db/migration

# Ensure wait-for.sh has execute permissions
RUN chmod +x ./wait-for.sh ./start.sh

EXPOSE 8080 8090
CMD [ "services/manage-service/main" ]
ENTRYPOINT [ "/app/start.sh" ]