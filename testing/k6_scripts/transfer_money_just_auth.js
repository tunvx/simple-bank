import http from 'k6/http';
import { check, sleep } from 'k6';
import { Counter } from 'k6/metrics';
import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js'

export const options = {
    vus: __ENV.VUS ? parseInt(__ENV.VUS) : 1,
    iterations: __ENV.ITER ? parseInt(__ENV.ITER) : undefined,
    duration: __ENV.DURATION ? __ENV.DURATION : undefined,
};


let totalAccounts = 50000;
// let apiHost = "http://localhost:8082";                      // Docker 
let apiHost = "http://transaction-service.banking.com";     // Minikube
let failedRequestCounter = new Counter('failed_requests');

const BEARER_TOKEN = "v2.public.eyJpZCI6IjU0MzVmNDkyLWYxMDItNGIwNS1hMzZmLWFjZmMxMzI5MWM5NiIsInVzZXJfaWQiOjEsInJvbGUiOiJiYW5rZXIiLCJpc3N1ZWRfYXQiOiIyMDI0LTEyLTAzVDEyOjExOjA5Ljk4MzQwOTgzNloiLCJleHBpcmVkX2F0IjoiMjAyNC0xMi0wNFQxMjoxMTowOS45ODM0MDk5NjFaIn2zqhR2ZLrR9_gbaqUl704kgHNXFe5ZyUtNQVX5TF_j_zox_WeF8-5QN17Xd9igW9MR7xkAJXhl_GTe8PVMabwI.bnVsbA";

export default function () {
    // Generate two distinct random IDs between 1 and totalAccounts
    let senderId = Math.floor(Math.random() * totalAccounts) + 1;
    let recipientId;
    do {
        recipientId = Math.floor(Math.random() * totalAccounts) + 1;
    } while (recipientId === senderId);

    let senderAccountNumber = `${senderId.toString().padStart(11, '0')}`;
    let recipientAccountNumber = `${recipientId.toString().padStart(11, '0')}`;

    // Perform internal transfer
    let transferUrl = `${apiHost}/v1/test/fast_internal_transfer_process_auth`;
    let transferPayload = JSON.stringify({
        amount: 10000,
        sender_acc_number: senderAccountNumber,
        recipient_bank_code: "Ngan Hang VCB",
        recipient_acc_number: recipientAccountNumber,
        recipient_name: "Nguyen Van CBA",
        currency_type: "VND",
        message: "Nguyen Van ABC chuyen tien"
    });

    let res = http.post(transferUrl, transferPayload, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${BEARER_TOKEN}`,
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

    const filename = `json_results/transfer_money_just_auth_VU_${options.vus}_${filenamePrefix}.json`;

    return {
        [filename]: JSON.stringify(data, null, 4),
    };
}