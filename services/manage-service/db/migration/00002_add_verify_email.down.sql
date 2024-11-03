-- Remove Foreign Key Constraint from verify_emails
ALTER TABLE verify_emails 
DROP CONSTRAINT emails_customers_customer_id_fkey;
-- Drop Verify Email Table
DROP TABLE verify_emails;

ALTER TABLE customers DROP COLUMN is_email_verified;