CREATE TABLE verify_emails (
  id uuid PRIMARY KEY,
  customer_id bigint NOT NULL,
  email_address varchar NOT NULL,
  secret_code varchar NOT NULL,
  is_used bool NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT (now()),
  expired_at timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE verify_emails 
ADD FOREIGN KEY (customer_id) REFERENCES customers (customer_id);

ALTER TABLE customers ADD COLUMN is_email_verified bool NOT NULL DEFAULT false;