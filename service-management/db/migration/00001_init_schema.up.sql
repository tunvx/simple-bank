
-- 1. Status of transactions in the banking system
CREATE TYPE TransactionStatus AS ENUM (
  'completed',  -- Transaction is completed
  'pending',    -- Transaction is pending
  'failed'      -- Transaction has failed
);


-- 2. Types of transactions in the banking system
CREATE TYPE TransactionType AS ENUM (
  'internal_send',      -- Internal transfer money
  'internal_receive',   -- Internal receive money
  'external_send',      -- External transfer money
  'external_receive',   -- External receive money
  'repay_loan',         -- Repay for your loan
  'deposit_savings',    -- Deposis from your savings
  'other'               -- Other transactions
);

-- 3. Supported currency types in the banking system
CREATE TYPE CurrencyType AS ENUM (
  'VND',  -- Vietnamese Dong
  'USD'   -- United States Dollar
);

-- 4. Status of bank accounts
CREATE TYPE AccountStatus AS ENUM (
  'active',   -- Account is active
  'inactive'  -- Account is inactive or locked
);

-- 5. Customer tiers based on value and priority
CREATE TYPE CustomerTier AS ENUM (
  'regular',   
  'bronze',     
  'silver',     
  'gold',       
  'platinum',   
  'diamond'     
);

-- 6. Customer segments in the banking system
CREATE TYPE CustomerSegment AS ENUM (
  'individual',          
  'small_enterprise',    
  'medium_enterprise',   
  'large_enterprise',   
  'institutional'        
);

-- 7. Financial status of customer
CREATE TYPE FinancialStatus AS ENUM (
  'excellent',   -- Excellent, with a string credit history and no bad debts
  'very_good',   -- Very good, with stable assers and income, minimal debt
  'good',        -- Good, with stable credit history, but may have some debt
  'fair',        -- Fair, with some credit issues or debt
  'poor',        -- Poor, with bad debts or unstable income
  'very_poor'    -- Very poor, with severe financial difficulties or high bad debt
);



-- Bảng khách hàng của ngân hàng
CREATE TABLE customers (
  customer_id bigint PRIMARY KEY DEFAULT shard_id_generator.generate_id(),
  customer_rid varchar(15) UNIQUE NOT NULL,
  fullname varchar NOT NULL,
  date_of_birth date NOT NULL,
  address varchar NOT NULL,
  phone_number varchar(15) UNIQUE NOT NULL,
  email varchar UNIQUE NOT NULL,
  customer_tier CustomerTier NOT NULL,
  customer_segment CustomerSegment NOT NULL,
  financial_status FinancialStatus NOT NULL
);

-- Bảng tài khoản của khách hàng
CREATE TABLE accounts (
  account_id bigint PRIMARY KEY DEFAULT shard_id_generator.generate_id(),
  account_number varchar(15) UNIQUE NOT NULL,
  customer_id bigint NOT NULL,
  current_balance bigint NOT NULL,
  currency_type CurrencyType NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  description text NOT NULL DEFAULT '...',
  account_status AccountStatus NOT NULL
);

-- Bảng giao dịch chuyển tiền theo tài khoản khách hàng
CREATE TABLE money_transfer_transactions (
  transaction_id bigint PRIMARY KEY DEFAULT shard_id_generator.generate_id(),
  amount bigint NOT NULL,
  account_id bigint NOT NULL,
  new_balance bigint NOT NULL,
  transaction_time timestamptz NOT NULL DEFAULT (now()),
  description text NOT NULL DEFAULT '...',
  transaction_type TransactionType NOT NULL,
  transaction_status TransactionStatus NOT NULL
);

CREATE INDEX ON customers (customer_segment);

CREATE INDEX ON customers (customer_tier);

CREATE INDEX ON customers (financial_status);

CREATE INDEX ON accounts (customer_id);

CREATE INDEX ON accounts (account_status);

CREATE INDEX ON money_transfer_transactions (account_id);

ALTER TABLE accounts ADD FOREIGN KEY (customer_id) REFERENCES customers (customer_id);

ALTER TABLE money_transfer_transactions ADD FOREIGN KEY (account_id) REFERENCES accounts (account_id);
