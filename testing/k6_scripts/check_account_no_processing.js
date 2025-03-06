import http from 'k6/http';
import { check, sleep } from 'k6';
import { Counter } from 'k6/metrics';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js'

export const options = {
    vus: __ENV.VUS ? parseInt(__ENV.VUS) : 1,
    iterations: __ENV.ITER ? parseInt(__ENV.ITER) : undefined,
    duration: __ENV.DURATION ? __ENV.DURATION : undefined,
};


let totalAccounts = 20000;
let apiHost = "http://localhost:8083";               // Docker 
// let apiHost = "http://moneytransfer.banking.local";     // Minikube
let failedRequestCounter = new Counter('failed_requests');

export default function () {
    let checkerId = Math.floor(Math.random() * totalAccounts) + 1;
    let accountNumber = `${checkerId.toString().padStart(11, '0')}`;

    let checkAccountUrl = `${apiHost}/v1/test/check_account_no_processing`;
    let checkAccountPayload = JSON.stringify({
        acc_number: accountNumber,
        currency_type: "VND",
    });

    let res = http.post(checkAccountUrl, checkAccountPayload, {
        headers: {
            'Content-Type': 'application/json',
            // 'Authorization': `Bearer ${BEARER_TOKEN}`,
        },
    });

    let isSuccess = check(res, {
        'is status 200': (r) => r.status === 200,
    });

    if (!isSuccess) {
        console.error(`Request failed: ${res.status} ${res.body}`);
        failedRequestCounter.add(1);
        return;
    }
}

export function handleSummary(data) {
    console.log('Preparing the end-of-test summary...');

    // Print to the console for real-time feedback
    console.log(textSummary(data, { indent: ' ', enableColors: true }));

    // Determine the filename based on the test configuration
    let filenamePrefix = options.iterations 
        ? `ITER_${options.iterations}` 
        : `DURATION_${options.duration}`;

    // Convert the filename prefix to uppercase
    filenamePrefix = filenamePrefix.toUpperCase();

    const filename = `json_results/check_account_no_processing_VU_${options.vus}_${filenamePrefix}.json`;

    return {
        [filename]: JSON.stringify(data, null, 4),
    };
}