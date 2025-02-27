# Build stage
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN cd vnb-auth-service && go mod download
RUN cd vnb-auth-service && go build -o main main.go

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/vnb-auth-service/main .
COPY config.env .
COPY start.sh .
COPY wait-for.sh .
COPY vnb-auth-service/db/migration ./db/migration

# Ensure wait-for.sh has execute permissions
RUN chmod +x ./wait-for.sh ./start.sh

EXPOSE 80 443
CMD [ "vnb-auth-service/main" ]
ENTRYPOINT [ "/app/start.sh" ]