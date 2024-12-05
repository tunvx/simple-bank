import bcrypt
from datetime import datetime
import asyncpg
import asyncio

def hash_password(password: str) -> str:
    # Encode the password to bytes
    password_bytes = password.encode('utf-8')
    
    # Generate bcrypt hash of the password
    hashed_password = bcrypt.hashpw(password_bytes, bcrypt.gensalt())
    
    # Decode the hash to a string and return
    return hashed_password.decode('utf-8')

async def generate_50k_customers(n, db_url):
    # Connect to your PostgreSQL database
    conn = await asyncpg.connect(db_url)

    # Convert the date_of_birth string to a date object
    date_of_birth = datetime.strptime('2001-07-17', '%Y-%m-%d').date()

    # Insert records one by one
    for id in range(1, n):
        customer_rid = f"Rid-{str(id).zfill(8)}"
        fullname = f"Cus-{str(id).zfill(8)}"
        address = f"Add-{str(id).zfill(8)}"
        phone_number = f"0000-{str(id).zfill(8)}"
        email = f"example.{str(id).zfill(8)}@gmail.com"
        customer_tier = 'gold'
        customer_segment = 'individual'
        financial_status = 'good'
        
        # Execute the command with the date object
        await conn.execute("""
            INSERT INTO customers (customer_rid, fullname, date_of_birth, address, phone_number, email, customer_tier, customer_segment, financial_status)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        """, customer_rid, fullname, date_of_birth, address, phone_number, email, customer_tier, customer_segment, financial_status)
        
    # Close the connection
    await conn.close()
    print(f"{n-1} customer records have been inserted into the database.")


async def generate_50k_customer_credentials(n, core_db_url, auth_db_url):
    core_conn = await asyncpg.connect(core_db_url)
    auth_conn = await asyncpg.connect(auth_db_url)

    hashed_password = hash_password("2024@Aug")

    for id in range(1, n):
        customer_rid = f"Rid-{str(id).zfill(8)}"
        username = f"user{str(id).zfill(8)}"

        customer_id = await core_conn.fetchval("SELECT customer_id FROM customers WHERE customer_rid = $1", customer_rid)
        
        if customer_id:
            await auth_conn.execute("""
                INSERT INTO customer_credentials (customer_id, username, hashed_password)
                VALUES ($1, $2, $3)
            """, customer_id, username, hashed_password)
        else:
            print(f"Customer with RID {customer_rid} not found.")

    await core_conn.close()
    await auth_conn.close()
    print(f"{n-1} customer credentials records have been inserted into the database.")

async def generate_50k_account(n, db_url):
    conn = await asyncpg.connect(db_url)

    for id in range(1, n):
        customer_rid = f"Rid-{str(id).zfill(8)}"
        customer_id = await conn.fetchval("SELECT customer_id FROM customers WHERE customer_rid = $1", customer_rid)

        if customer_id:
            account_number = f"{str(id).zfill(11)}"
            current_balance = 500000000
            currency_type = "VND"
            description = "....."
            account_status = "active"

            await conn.execute("""
                INSERT INTO accounts (account_number, customer_id, current_balance, currency_type, description, account_status)
                VALUES ($1, $2, $3, $4, $5, $6)
            """, account_number, customer_id, current_balance, currency_type, description, account_status)
        else:
            print(f"Customer with RID {customer_rid} not found.")

    await conn.close()
    print(f"{n-1} account records have been inserted into the database.")

# Running the async tasks
async def main():
    n = 50001
    core_db_url = "postgresql://postgres:secret@localhost:5432/core_db?sslmode=disable"
    auth_db_url = "postgresql://postgres:secret@localhost:5433/auth_db?sslmode=disable"

    await generate_50k_customers(n, core_db_url)
    await generate_50k_customer_credentials(n, core_db_url, auth_db_url)
    await generate_50k_account(n, core_db_url)

# Entry point
if __name__ == "__main__":
    asyncio.run(main())