# Simple Bank - Microservices Architecture - Docker Deployment

This project implements the server-side of a simple banking system using a microservices architecture. It provides core banking functionalities such as customer and account management, secure authentication, money transfers, notifications, and monitoring capabilities.

Note: Loan and saving features are planned for future development.

### Banking System Architecture

![Architecture Diagram](./SimpleBank.png)

### Description of banking service
1. **Management Service**
  + **Customer Registration:** Allows new customers to register, with a background process to send email verification.
  + **Account Creation:** Allows customers to create their accounts.
2. **Authentication Service**
+ **Customer Credential Creation:** Allows customers to create login credentials.
+ **Login and Session Management:** Provides secure login functionality and session management.
3. **Transfer Money Service**
+ **Money Transfer Transactions:** Support secure, consistent and efficient money transfers between accounts or banks.
4. **Notification Service**
- **Email Verification:** Sends verification emails to customers as part of the registration process.
- **Transaction Notifications:** Sends email notifications to customers after each transaction.

## Tech Stack (for Docker deployment)

+ **Architecture:** Microservices, Service-Oriented Architecture (SOA).

+ **Languages:** Golang.

+ **DB:** PosgreSQL.

+ **APIs:** RESTful (client-to-service), gRPC ( service-to-service).

+ **Build/Test/Deployment:** Docker (docker file, docker build), Docker Compose, Unittest, K6 (performance testing).

+ **Advanced Techs:** JWT/Paseto (authentication/security), Kafka (message queue), Redis (caching), Logging and monitoring.

## Main Purpose:: Handling Large-Scale Money Transfers.
### Solution 1: Single PostgreSQL Instance (ACID for money transfer transactions).
- Limitations? Vertical scaling only (limited by hardware capacity), Single point of failure (SPOF).
- When to Use? Solution 1 is viable for moderate traffic (<50k RPS).

### Solution 2 (Improvement): Horizontal Scaling with DB Sharding (ACID per shard) + Eventual Consistency.
- When to Use? Solution 2 is essential for high-scale systems (200K+ RPS).

### Detailed Solution 2: I will simulate the process of money transfer transaction between two database instances (two shards) (A → B).

### 1. A (Source Account) Processes Deduction & Sets Status to PENDING
1. **Deduct funds from the source account** (`UPDATE accounts SET balance = balance - amount WHERE id = source_id`).
2. **Create a sending transaction record with status PENDING** (`INSERT INTO transactions ... with PENDING status`)
3. Send a transfer request to **B** (Destination Account).

### 2. B (Destination Account) Processes and Responds
1. Validate the destination account.
2. If valid:
    - **Add funds to destination account** (`UPDATE accounts SET balance = balance + amount WHERE id = destination_id`).
    - **Create a receiving transaction record with status SUCCESS** (`INSERT INTO transactions ... with SUCCESS status`)
    - Send an **OK** response to **A**.
3. If invalid:
   - Send a **FAILED** response to **A**.

### 3. A (Source Account) Updates Transaction Status
- If **OK** response received, update transaction status from **PENDING** to **SUCCESS**.
- If **FAILED** response received, refund the source account (**ROLLBACK**).
- If no response received, continue periodic retry attempts (retry mechanism).

### 4. Handling Errors or No Response from B
If multiple retry attempts fail, the transaction remains **PENDING**
  - Mark the transaction as **"manual review required"**.
  - Notify support team to investigate.

### Why Is This Design Reasonable?

✅ **Avoids distributed transactions (2PC)** → Simpler, no resource locking across databases.

✅ **Ensures eventual consistency** → Prevents fund loss even in network failures.

✅ **Has a recovery mechanism** → Retry logic or manual intervention ensures resolution.


## Docker Deployment - Quick Start Guide
### Docker Deployment Steps
1. Clone the project:
```bash
  git clone https://github.com/tunvx/simple-bank
```

2. Navigate to the project directory:
```bash
  cd simple-bank
```

3. Switch to the docker-deploy branch:
```bash
  git checkout docker-deploy
```

4. Navigate to the `domolo` directory (short for docker + monitoring + logging):
```bash
  cd domolo
```

5. Start core banking services (database, Redis, Kafka, etc.):
```bash
  make start-infra
```

6. Start the monitoring services (Prometheus, Grafana, etc.):
```bash
  make start-monitor
```

### Docker Monitoring Steps
After deploy monitor, do the following steps:
+ Login to grafana ( admin : abc13579 )
+ Connections -> Add a new connection -> Find and enter "Prometheus" -> Add a new data source -> Enter "http://prometheus:9090" into "Prometheus server URL" -> Save and Test
+ Dashboards -> New -> Import -> Enter "1860" and "193" ID (for node-exporter and cadvisor) -> Select data source is "prometheus" -> Import -> You can see defaul dashboards -> Save

### Testing Result (Performance)
**Load testing results with K6 (100 VU for 3 minutes)**

| No. | Framework | Test              | Duration (min) | Total Requests (req) | HTTP RPS (req/s)  | Avg Latency (ms) | Min Latency (ms) | Max Latency (ms) | P95 Latency (ms) |
|-----|-----------|-------------------|----------------|-----------------------|------------------|------------------|------------------|------------------|------------------|
| 1   | Raw       | Empty GET         | 3              | -                     | *                | -                | -                | -                | -                |
| 2   | Raw       | Empty POST        | 3              | -                     | *                | -                | -                | -                | -                |
| 3   | gRPC      | Empty GET         | 3              | 6,128,749             | 34,047           | 2.9              | 0.167            | 59.04            | 4.56             |
| 4   | gRPC      | Empty POST        | 3              | 5,344,435             | 29,690           | 3.32             | 0.194            | 34.86            | 6.4              |
| 5   | gRPC      | Check Account     | 3              | 3,251,939             | 18,065           | 5.49             | 0.329            | 127.54           | 11.77            |
| 6   | gRPC      | Transfer Money    | 3              | 1,278,048             | 7,099            | 14.0             | 2.07             | 588.27           | 22.95            |


Refer to the testing folder for details (experiment and analyze results):
```bash
  https://github.com/tunvx/simple-bank/tree/docker-deploy/testing
```

## Deploy to Kubernetes

## Documentation

## Appendix

### Techniques tags 
The tags I read this while doing this project. I make notes of them because I think it's useful to learn about
+ _ monolithic, microservice, SOA, distributed systems, golang.
+ _ design_database, db_nomalization, db_indexes, db_migration (sqlc), transaction, ACID, consistency_locking.
+ _ RESTful, gRPC, HTTP/1.1, HTTP/2, RPC, HTTPS, SSL/TLS.
+ _ token_based_authentication, JWT, Paseto, session_management, access_control.
+ _ unittest, performance_test, load_testing, k6.
+ _ containerization, docker, dockerfile, docker_compose, kubernetes (k8s).
+ _ redis, kafka, message_queue, background_worker, asynchronous_communication, asynchronous_processing, caching, logging, monitoring, alerting, metrics collection.


