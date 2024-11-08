# Simple Bank (in Microservices)

This project implements the server-side of a simple banking system. It leverages a microservices architecture to deliver a range of banking functionalities, including customer and account management, secure authentication, money transfer transactions, notifications, historical data management, logging, and monitoring.

Note: Loan and savings features are planned for future releases.

**Banking System Architecture:**

![Architecture Diagram](./SimpleBank.jpg)

### Description of banking service
1. **Manage Service**
  + Customer Registration: Allows new customers to register, with a background process to send email verification.
  + Account Creation: Allows customers to create thier accounts.
2. **Auth Service**
+ Customer Credential Creation: Allows customers to create login credentials.
+ Login and Session Management: Provides secure login functionality and session management.
3. **Transaction Service**
+ Money Transfer Transactions: Support secure, consistent and efficient money transfers between accounts.
4. **Notification Service**
- Email Verification: Sends verification emails to customers as part of the registration process.
- Transaction Notifications: Sends email notifications to customers after each transaction.

## Tech Stack

+ **Architecture:** Microservices, Service-Oriented Architecture (SOA).

+ **Languages:** Golang.

+ **DB:** PosgreSQL.

+ **APIs:** RESTful (client-to-service), gRPC ( service-to-service).

+ **Build/Test/Deployment:** CI/CD with GitHub Actions, unit and performance testing, Docker Compose.

+ **Advanced Techs:** JWT/Paseto (authentication/security), Kafka (message queue), Redis (caching), logging, monitoring.

+ **Planned:** Distributed SQL, Docker Swarm/Kubernetes for auto-scaling.

## Setup local development (MacOS)
### Install tools
+ [Docker desktop](https://www.docker.com/products/docker-desktop)
+ [TablePlus](https://tableplus.com/)
+ [Homebrew](https://brew.sh/)
+ [Golang 1.22](https://golang.org/)
+ Sqlc
```bash
  brew install sqlc
```
+ Protoc
```bash
  brew install protobuf
```
+ Migrate
```bash
  brew install golang-migrate
```

## Deploy as Docker Containers
+ Clone the project
```bash
  git clone https://github.com/tunvx/simple-bank
```

+ Go to the project directory
```bash
  cd simple-bank
```

+ Go to the domolo directory (docker + monitoring + logging) 
```bash
  cd domolo
```

+ Start the server
```bash
  make start-infra
```

+ Start the monitor
```bash
  make start-monitor
```

### Steps for a basic monitoring
After deploy monitor, do the following steps:
+ Login to grafana ( admin : abc13579 )
+ Connections -> Add a new connection -> Find and enter "Prometheus" -> Add a new data source -> Enter "http://prometheus:9090" into "Prometheus server URL" -> Save and Test
+ Dashboards -> New -> Import -> Enter "1860" and "193" ID (for node-exporter and cadvisor) -> Select data source is "prometheus" -> Import -> You can see defaul dashboards -> Save

## Testing & Performance
+ Refer to the testing folder for details.

## Appendix

### Techniques tags 
The tags I read this while doing this project. I make notes of them because I think it's useful to learn about
+ _ monolithic, microservice, SOA, distributed systems, golang.
+ _ design_database, db_nomalization, db_indexes, db_migration (sqlc), transaction, ACID, consistency_locking.
+ _ RESTful, gRPC, HTTP/1.1, HTTP/2, RPC, HTTPS, SSL/TLS.
+ _ token_based_authentication, JWT, Paseto, session_management, access_control.
+ _ unittest, performance_test, load_testing, k6.
+ _ docker, docker_compose, dockerfile, containerization.
+ _ redis, kafka, message_queue, background_worker, asynchronous_communication, asynchronous_processing, caching, logging, monitoring, alerting, metrics collection.

