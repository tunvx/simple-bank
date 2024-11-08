#!/bin/bash

vus=100
duration="3m"

scripts=(
  "k6_scripts/check_account_just_auth.js"
  "k6_scripts/check_account_no_processing.js"
  "k6_scripts/check_account.js"
  "k6_scripts/empty_get.js"
  "k6_scripts/empty_post.js"
  "k6_scripts/transfer_money_just_auth.js"
  "k6_scripts/transfer_money_no_processing.js"
  "k6_scripts/transfer_money.js"
)

sleep_time=180

for script in "${scripts[@]}"; do
  echo "Running script: $script"
  k6 run --vus $vus --duration $duration $script
  echo "Completed script: $script. Sleeping for $sleep_time seconds."
  sleep $sleep_time
done

echo "All scripts completed."
