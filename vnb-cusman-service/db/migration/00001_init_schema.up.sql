
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
  'standard',    -- Khách hàng tiêu chuẩn (thấp nhất)
  'bronze',      -- Hạng đồng
  'silver',      -- Hạng bạc
  'gold',        -- Hạng vàng
  'platinum',    -- Hạng bạch kim
  'diamond',     -- Hạng kim cương (cao nhất)
  'vip'          -- Khách hàng VIP đặc biệt
);

-- 6. Customer segments in the banking system
CREATE TYPE CustomerSegment AS ENUM (
  'retail',           -- Cá nhân (khách hàng cá nhân, hộ gia đình)
  'small_business',   -- Doanh nghiệp nhỏ
  'corporate',        -- Doanh nghiệp lớn (công ty vừa và lớn)
  'institutional',    -- Tổ chức tài chính (quỹ, tổ chức lớn)
  'government'        -- Cơ quan chính phủ, nhà nước
);

-- 7. Financial status of customer
CREATE TYPE FinancialStatus AS ENUM (
  'very_good',    -- Rất tốt (tài sản ổn định, không nợ xấu)
  'good',         -- Tốt (tín dụng ổn định, có thể có một ít nợ)
  'average',      -- Trung bình (một số vấn đề tín dụng nhỏ)
  'low_risk',     -- Rủi ro thấp (có thể có nợ nhưng kiểm soát được)
  'high_risk',    -- Rủi ro cao (nợ xấu hoặc thu nhập không ổn định)
  'defaulted'     -- Mất khả năng thanh toán (nợ xấu nghiêm trọng)
);


-- Bảng khách hàng của ngân hàng
CREATE TABLE customers (
  customer_id bigint PRIMARY KEY,
  customer_rid varchar(15) UNIQUE NOT NULL,
  full_name varchar NOT NULL,
  date_of_birth date NOT NULL,
  address varchar NOT NULL,
  phone_number varchar(15) UNIQUE NOT NULL,
  email varchar UNIQUE NOT NULL,
  customer_tier CustomerTier NOT NULL,
  customer_segment CustomerSegment NOT NULL,
  financial_status FinancialStatus NOT NULL
);

CREATE INDEX ON customers (customer_segment);
CREATE INDEX ON customers (customer_tier);
CREATE INDEX ON customers (financial_status);


-- Bảng tài khoản của khách hàng
CREATE TABLE accounts (
  account_id bigint PRIMARY KEY,
  customer_id bigint NOT NULL,
  current_balance bigint NOT NULL,
  currency_type CurrencyType NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  description text NOT NULL DEFAULT '...',
  account_status AccountStatus NOT NULL
);

CREATE INDEX ON accounts (customer_id);
CREATE INDEX ON accounts (account_status);

ALTER TABLE accounts 
ADD FOREIGN KEY (customer_id) REFERENCES customers (customer_id);


-- Bảng giao dịch chuyển tiền theo tài khoản khách hàng
CREATE TABLE account_transactions (
  transaction_id uuid PRIMARY KEY,
  amount bigint NOT NULL,
  account_id bigint NOT NULL,
  new_balance bigint NOT NULL,
  transaction_time timestamptz NOT NULL DEFAULT (now()),
  description text NOT NULL DEFAULT '...',
  transaction_type TransactionType NOT NULL,
  transaction_status TransactionStatus NOT NULL
);

CREATE INDEX ON account_transactions (account_id);

ALTER TABLE account_transactions 
ADD FOREIGN KEY (account_id) REFERENCES accounts (account_id);
