CREATE SCHEMA IF NOT EXISTS shard_1;

CREATE SEQUENCE IF NOT EXISTS shard_1.global_id_sequence;

CREATE OR REPLACE FUNCTION shard_1.id_generator(OUT result bigint) AS $$
DECLARE
    our_epoch bigint := 1314220021721;
    seq_id bigint;
    now_millis bigint;
    shard_id int := 1;
BEGIN
    SELECT nextval('shard_1.global_id_sequence') % 1024 INTO seq_id;
    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
    result := (now_millis - our_epoch) << 23;
    result := result | (shard_id << 10);
    result := result | (seq_id);
END;
$$ LANGUAGE PLPGSQL; 

-- Status of transactions in the banking system
CREATE TYPE TransactionStatus AS ENUM (
  'pending',    -- Transaction is pending
  'completed',  -- Transaction is completed
  'failed'      -- Transaction has failed
);

-- Types of transaction in the banking system
CREATE TYPE TransactionType AS ENUM (
  'send_money_internal',       -- Internal transfer money
  'receive_money_internal',    -- Internal receive money
  'others'                     -- Other transactions
);

-- Supported currency types in the banking system
CREATE TYPE CurrencyType AS ENUM (
  'VND',  -- Vietnamese Dong
  'USD'   -- United States Dollar
);

-- Status of bank accounts
CREATE TYPE AccountStatus AS ENUM (
  'active',   -- Account is active
  'inactive'  -- Account is inactive or locked
);

-- Customer tiers based on value and priority
CREATE TYPE CustomerTier AS ENUM (
  'regular',   
  'bronze',     
  'silver',     
  'gold',       
  'platinum',   
  'diamond'     
);

-- Customer segments in the banking system
CREATE TYPE CustomerSegment AS ENUM (
  'individual',          
  'small_enterprise',    
  'medium_enterprise',   
  'large_enterprise',   
  'institutional'        
);

-- Financial status of custoemr
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
  customer_id bigint PRIMARY KEY DEFAULT shard_1.id_generator(),
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
  account_id bigint PRIMARY KEY DEFAULT shard_1.id_generator(),
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
  transaction_id bigint PRIMARY KEY DEFAULT shard_1.id_generator(),
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
