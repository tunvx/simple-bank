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
let apiHost = "http://localhost:8082";
let failedRequestCounter = new Counter('failed_requests');

const BEARER_TOKEN = "v2.public.eyJpZCI6IjNiMmE4NjhjLTlkYTYtNDg4ZS04MGNmLWZjMjllOGZiNWFhNyIsInVzZXJfaWQiOjEsInJvbGUiOiJiYW5rZXIiLCJpc3N1ZWRfYXQiOiIyMDI0LTExLTE4VDE0OjE4OjE0LjI3MzEwNTYyOFoiLCJleHBpcmVkX2F0IjoiMjAyNC0xMS0xOVQxNDoxODoxNC4yNzMxMDU3MTJaIn0ju6XaQSd8LYkGfaA9IZkIQC_aZwbzHYyLmgBT4ebpgSbWVq9Ij0Jc8eK8s_PalMAELkCGU3ZKPS3_KYHH__4A.bnVsbA";

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
    let transferUrl = `${apiHost}/v1/test/fast_internal_transfer_without_processing`;
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

    const filename = `json_results/transfer_money_no_processing_VU_${options.vus}_${filenamePrefix}.json`;

    return {
        [filename]: JSON.stringify(data, null, 4),
    };
}