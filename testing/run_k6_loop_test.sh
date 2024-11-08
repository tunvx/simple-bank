#!/bin/bash

k6_script=""
iterations=""

while [[ $# -gt 0 ]]; do
    case $1 in
        --file)
        k6_script="$2"
        shift 2
        ;;
        --iter|--iterations)
        iterations="$2"
        shift 2
        ;;
        *)
        echo "Unknown parameter: $1"
        exit 1
        ;;
    esac
done

if [[ -z "$k6_script" || -z "$iterations" ]]; then
    echo "Usage: ./run_k6_tests.sh --file <k6_script_path> --iterations <iterations>"
    exit 1
fi

# Define the list of VUs
vus_list=(10 20 30 40 50 75 100 125 150 175 200)

for vus in "${vus_list[@]}"
do
    docker exec -it postgres01 psql -U root -d core_db -c "TRUNCATE TABLE money_transfer_transactions;"
    if [[ $? -ne 0 ]]; then
        echo "Failed to truncate money_transfer_transactions table."
        exit 1
    fi

    echo "Running k6 with $vus VUs and $iterations iterations..."
    
    k6 run --vus $vus --iterations "$iterations" "$k6_script"
    
    echo "Waiting for 180 seconds before next run..."
    sleep 180
done

echo "All tests completed."
