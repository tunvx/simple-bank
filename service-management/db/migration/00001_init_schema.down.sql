-- Drop tables that reference ENUM types
DROP TABLE IF EXISTS money_transfer_transactions;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS customers;

-- Drop ENUM types
DROP TYPE IF EXISTS TransactionType;
DROP TYPE IF EXISTS TransactionStatus;
DROP TYPE IF EXISTS CurrencyType;
DROP TYPE IF EXISTS AccountStatus;
DROP TYPE IF EXISTS CustomerTier;
DROP TYPE IF EXISTS CustomerSegment;
DROP TYPE IF EXISTS FinancialStatus;