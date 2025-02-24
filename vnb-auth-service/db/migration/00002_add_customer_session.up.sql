CREATE TABLE customer_sessions (
  session_id uuid PRIMARY KEY,
  customer_id bigint NOT NULL,
  refresh_token varchar NOT NULL,
  user_agent varchar NOT NULL,
  client_ip varchar NOT NULL,
  is_blocked boolean NOT NULL DEFAULT false,
  expires_at timestamptz NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

-- Add Foreign Key Constraint to customer_sessions
ALTER TABLE customer_sessions 
ADD FOREIGN KEY (customer_id) REFERENCES customer_credentials (customer_id);