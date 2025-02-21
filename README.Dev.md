# System Design Analysis

# High Read & High Write Requirements

## High Read (normal actions + reporting actions)
1. look up (account's **shard_id**) by account_id .
2. look up (account info) by account_id.
3. look up (login info) by user_name, password.
4. look up (account transaction history) by account_id.
5. look up (data) for reporting.

## High write (normal actions)
1. update account's balance + transaction history
2. update savings + transaction history
3. update loans + transaction history

## Move data across shards
**Constrains:** All data of one user data must be in one shard

1. does transaction_id include shard_id?
    + yes -> 
    + no -> 
    2. does account_id include shard_id?
        + yes -> 
        + no -> 
        3. does customer_id include shard_id?
            + yes -> 
            + no -> 

# Database Structure Design
The system consists of three main databases:

## Lookup db 
+ customer_shard_map (required) 
+ account_shard_map (required)
+ transaction_shard_map (optinal)

(high read is here)

## Auth db
+ customer credential (contain shard_id)
+ customer sessions (UUID)

(high read is here)

## Core db
+ customers (pk: customer id)
+ accounts (pk: account id)
+ money_transfer_transaction (pk: transaction_id)
+ .....

(high write is here)

# API Description

## Authentication Service
Handles user authentication and credential management.

#### 1. Create Customer Credentials (`POST /auth/customers/credentials`)
- **Description:** Creates a new authentication record for a customer.
- **Input:**  `customer_id`, `username`, `password`
- **Output:**  `is_created` (bool)

#### 2. Customer Login (`POST /auth/login`)
- **Description:** Authenticates a customer and generates access tokens.
- **Input:**  `username`, `password`
- **Output:**  `session_id`, `access_token`, `refresh_token`, `access_token_expires_at`, `refresh_token_expires_at`

#### 3. Update Customer Credentials (`PUT /auth/customers/credentials`)
- **Description:** Updates the username and/or password of a customer.
- **Input:**  `customer_id`, `username`, `password`
- **Output:**  `is_updated` (bool)

#### 4. Update Customer Shard (`PUT /auth/customers/shard`)
- **Description:** Assigns or updates the shard ID for a customer (used in distributed database setups).
- **Input:**  `customer_id`, `shard_id`
- **Output:**  `is_updated` (bool)

---

## Customer Management Service
Handles customer lifecycle management, including account creation and verification.

#### 1. Create New Customer (`POST /customers`)
- **Description:** Creates a new customer record in the system.
- **Input:**  ...
- **Output:**  ...

#### 2. Verify Customer Email (`POST /customers/email/verify`)
- **Description:** Verifies a customer's email using a secret code.
- **Input:**  `email_address`, `secret_code`
- **Output:**  `is_verified` (bool)

#### 3. Create New Bank Account (`POST /accounts`)
- **Description:** Opens a new bank account for a customer.
- **Input:**  `account_id`, `currency_type`
- **Output:**  ...

#### 4. Get Customer Information (`GET /customers/{customer_id}`)
- **Description:** Retrieves details of a customer by ID.

#### 5. Get Account Information (`GET /accounts/{account_id}`)
- **Description:** Retrieves account details, including balance and status.

---

## Shard Management Service
Handles distributed database shard lookups.
#### 1. Lookup Account Shard (`GET /shards/accounts/{account_id}`)
- **Description:** Retrieves the shard ID assigned to a specific bank account.
- **Input:**  `account_id`
- **Output:**  `account_id`, `shard_id`

#### 2. Insert Account_id Into Shard (`POST /shards/accounts`)
- **Description:** ...
- **Input:**  `account_id`, `customer_id`, `shard_id`
- **Output:**  `is_inserted` (bool)

#### 3. Insert Customer_id Into Shard (`POST /shards/customers`)
- **Description:** ...
- **Input:**  `customer_rid`, `full_name`, `shard_id`
- **Output:**  `is_inserted` (bool), `customer_id`

#### 4. Lookup Customer Shard (`GET /shards/customers/{customer_rid}`)
- **Description:** Retrieves the shard ID assigned to a specific customer.
- **Input:**  `customer_rid`
- **Output:**  `customer_id`, `shard_id`

---

## Transfer Money Service
Handles internal and external money transfers.

### Money Receiving Process
#### 1. Verify Account (`POST /accounts/verify`)
- **Description:** Checks if the recipient account exists and is active.
- **Input:**  `account_id`, `currency_type`
- **Output:**  `is_valid`, `account_id`, `account_holder_name`

#### 2. Credit Funds (`POST /transactions/credit`) 
- **Description:** Adds money to the recipient's account.
- **Input:**  `amount`, `currency_type`, `to_account_id`, `transaction_id`, `description`
- **Output:**  `status`, `message`

---

### Internal Money Transfer Process (Within the Same Bank)
#### 1. Verify Account (`POST /accounts/verify`)
- **Description:** Checks if the sender and recipient accounts exist and are valid.
- **Input:**  `account_id`, `currency_type`
- **Output:**  `is_valid`, `account_id`, `account_holder_name`

#### 2. Internal Transfer (`POST /transactions/internal-transfer`)
- **Description:** Transfers funds between accounts within the same bank.
- **Input:**  `amount`, `currency_type`, `from_account_id`, `to_account_id`, `transaction_id`, `description`
- **Sub-process:** Calls `Credit Funds` (`POST /transactions/credit`)
- **Output:**  `status`, `message`

---

### External Money Transfer Process (Interbank Transfer)
#### 1. Verify External Account (`POST /partners/accounts/verify`)
- **Description:** Calls the partner bank’s API to check if the recipient account exists.

#### 2. External Transfer (`POST /transactions/external-transfer`)
- **Description:** Initiates a money transfer to an external bank.
- **Input:**  `amount`, `currency_type`, `from_account_id`, `to_account_id`, `to_bank_id`, `transaction_id`, `description`
- **Sub-process:** Calls `Credit Funds` API of the partner bank (`POST /partners/transactions/credit`)
- **Output:**  `status`, `message`