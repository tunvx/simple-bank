-- Create Customer Shard Map Table
CREATE TABLE customer_shard_map (
  customer_id BIGSERIAL PRIMARY KEY,
  customer_rid varchar UNIQUE NOT NULL,
  shard_id int NOT NULL
);

CREATE INDEX ON customer_shard_map (shard_id);

-- Create Account Shard Map Table
CREATE TABLE account_shard_map (
  account_id bigint PRIMARY KEY,
  customer_id bigint NOT NULL
);

CREATE INDEX ON account_shard_map (customer_id);

-- Add Foreign Key Constraint to account_shard_map
ALTER TABLE account_shard_map 
ADD FOREIGN KEY (customer_id) REFERENCES customer_shard_map (customer_id);