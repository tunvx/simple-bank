CREATE TABLE verify_emails (
  id bigint PRIMARY KEY DEFAULT shard_1.id_generator(),
  customer_id bigint NOT NULL,
  email varchar NOT NULL,
  secret_code varchar NOT NULL,
  is_used bool NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT (now()),
  expired_at timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

-- Add Foreign Key Constraint to customer_sessions
ALTER TABLE verify_emails 
ADD CONSTRAINT emails_customers_customer_id_fkey 
FOREIGN KEY (customer_id) REFERENCES customers (customer_id);

ALTER TABLE customers ADD COLUMN is_email_verified bool NOT NULL DEFAULT false;