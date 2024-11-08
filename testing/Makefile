mockdata:
	python3 mockdata/generate_mockdata.py

# k6maxconn:
# 	k6 run load_testing/max_conn.js

# k6checkaccount:
# 	k6 run load_testing/check_account.js

# k6transfermoney:
# 	k6 run load_testing/transfer_money.js

# k6loadtesting:
# 	k6 run load_testing/load_testing.js

.PHONY: mockdata

# TRUNCATE TABLE money_transfer_transactions;
# docker exec -it postgres01 psql -U root -d core_db -c "TRUNCATE TABLE money_transfer_transactions;"
