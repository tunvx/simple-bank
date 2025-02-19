# Build stage
FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN cd service-notification && go mod download
RUN cd service-notification && go build -o main main.go

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/service-notification/main .
COPY config.env .
COPY start.sh .
COPY wait-for.sh .

# Ensure wait-for.sh has execute permissions
RUN chmod +x ./wait-for.sh ./start.sh

EXPOSE 80 443
CMD [ "service-notification/main" ]
ENTRYPOINT [ "/app/start.sh" ]