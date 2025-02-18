-- Create Customer Credential Table
CREATE TABLE customer_credentials (
  customer_id bigint PRIMARY KEY,
  username varchar(15) UNIQUE NOT NULL,
  hashed_password varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now()),
  username_changed_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'), 
  password_changed_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);
