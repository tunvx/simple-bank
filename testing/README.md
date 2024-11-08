# Testing for Simple Bank


## 1. Load Testing Overview

The load testing is set up to assess the performance of the `transfer_money` operation under heavy load conditions. By running load tests with k6, I simulate concurrent transfers to evaluate the system's response times, throughput, and resource utilization.

## 2. Prerequisites

- [k6](https://k6.io/) - The load testing tool used for this project.
- Python 3 - For generating mock data needed during tests.
- Install the necessary Python packages
    ````bash
    pip3 install bcrypt asyncpg asyncio
    ````

## 3. Running the Load Test
1. Ensure you're in the testing directory
    ````bash
    cd testing
    ````

2. Generate mock data (needed for the first time):
    ```bash
    # Generate 50000 customers, credential, accounts (see code)
    make mockdata
    ```

3.  Generate a new test token (due to Expired token)
+ Before running the tests, generate a new test token by calling the test token endpoint (**after start-infra**) and replace the BEARER_TOKEN in test scripts with the newly generated token.  Use the following HTTP request to get a new token:
    + **URL:** http://127.0.0.1:8081/v1/generate_test_access_token

4. **Run the load test with:**

+ **Run Individual Tests:**
    ```bash
    # Run each test script with 100 Virtual Users (VUs) for a duration of 3 minutes. 
    # Use the following commands:

    # Empty GET request
    k6 run --vus 100 --duration 3m k6_scripts/empty_get.js

    # Empty POST request
    k6 run --vus 100 --duration 3m k6_scripts/empty_post.js

    # Check account with no processing
    k6 run --vus 100 --duration 3m k6_scripts/check_account_no_processing.js

    # Transfer money with no processing
    k6 run --vus 100 --duration 3m k6_scripts/transfer_money_no_processing.js

    # Check account just process authentication
    k6 run --vus 100 --duration 3m k6_scripts/check_account_just_auth.js

    # Transfer money just process authentication
    k6 run --vus 100 --duration 3m k6_scripts/transfer_money_just_auth.js

    # Check account
    k6 run --vus 100 --duration 3m k6_scripts/check_account.js

    # Transfer money
    k6 run --vus 100 --duration 3m k6_scripts/transfer_money.js
    ```

+ **Run All Script Tests:**
    ```bash
    # Params: fixed -> (100 VUs, duration 3 minutes) 
    ./run_k6_all_test.sh    
    ```

+ **Run Loop For One Script Test:**
    ```bash
    # Params: `--file` and `--iter`
    ./run_k6_loop_test.sh --file k6_scripts/transfer_money.js --iter 1000000
    ```

## 4. Test Experiment and Result
### 4.1. Architecture and Resource
**Test Machine:**
+ Model: MacBook Pro M1 (2021)
+ OS: macOS Sequoia 15.0.1
+ CPU: 10-core, 10-thread M1 Pro chip
+ Memory: 16GB RAM
+ Storage: 512GB SSD

**System Architecture:**
+ Deployment Model: Single container per service.

**Environment Details (Docker Virtual Machine):**
+ CPU: All 10-core.
+ Memory: 12GB RAM.
+ Storage: 256GB SSD.

### 4.2. Result

**Result on K6 with 100 VUs for 3m duration iterations:**

| No. | Framework | Test              | Duration (min) | Total Requests (req) | HTTP RPS (req/s)  | Avg Latency (ms) | Min Latency (ms) | Max Latency (ms) | P95 Latency (ms) |
|-----|-----------|-------------------|----------------|-----------------------|------------------|------------------|------------------|------------------|------------------|
| 1   | Raw       | Empty GET         | 3              | -                     | *                | -                | -                | -                | -                |
| 2   | Raw       | Empty POST        | 3              | -                     | *                | -                | -                | -                | -                |
| 3   | gRPC      | Empty GET         | 3              | 6,128,749             | 34,047           | 2.9              | 0.167            | 59.04            | 4.56             |
| 4   | gRPC      | Empty POST        | 3              | 5,344,435             | 29,690           | 3.32             | 0.194            | 34.86            | 6.4              |
| 5   | gRPC      | Check Account     | 3              | 3,251,939             | 18,065           | 5.49             | 0.329            | 127.54           | 11.77            |
| 6   | gRPC      | Transfer Money    | 3              | 1,278,048             | 7,099            | 14.0             | 2.07             | 588.27           | 22.95            |


For additional details of k6 metrics, refer to the [k6 documentation](https://grafana.com/docs/k6/latest/using-k6/metrics/reference/).

### 4.3 Limitations
For more details, refer to:

+ **Folder 1:** archived results run-all-01
+ **Folder 2:** archived results run-loop-01

**Preliminary Conclusion:**

The tests hit the 12GB memory limit of the Docker Desktop virtual machine (  image `archived results run-all-01/03 Node-Exporter-Full.png`).

And during the transfer money test, the Docker Desktop virtual machine likely hit its IO limit, as IO utilization peaked at approximately 100% (see image `archived results run-loop-01/03 Node-Exporter-Full.png`).

### 4.4. Configurations
#### 4.4.1. Configuration Number of Connection Pool
**TL;DR: Keep the default.**

**Increasing connections in the pool only helps when the time it takes to run a request is long**. To estimate a reasonable pool size, the following formula can be used:

```
    number_of_connection = (RPS * avg_request_latency_in_seconds) + add_on_factor
```

In the **simplest architecture** (one container per service), the following connection pool configurations were applied:
- **Transaction-service <-> Database**: Connection pool size ranging from 50 to 150.
- **Transaction-service <-> Redis Cache**: Maximum connection pool size of 120.
- Other services retained the default connection pool settings.

Note: Both the **manage-service** and **transaction-service** connect to the **core-database**. --> Change the number of core database pool connections.

However, results from the test experiments showed that modifying the connection pool size did not significantly impact performance. For simplicity and reliability, the default configuration is sufficient.

#### 4.4.2. Caching Strategy
**TL;DR: Cache all accounts.**

In the test experiments with **50,000 accounts**, all accounts were cached with a **TTL (Time-To-Live)** set to exceed the total duration of the experiment (a simple caching strategy because fields for the checking task rarely change in practice). It means when an update is needed, update the cache first.

For account checking tasks:

+ The first access retrieves data from the database for each account (**50,000 out of 1,000,000 total requests**).
+ This results in a cache hit ratio exceeding 95% (minimum), as subsequent requests for the same accounts are served directly from the cache.

## 5. Questions
+ How to optimize banking system? :)) 

