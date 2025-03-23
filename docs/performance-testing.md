# 1. Test Summary
| No. | Test Scenario | Resource | Duration (min) | TPS (tran/s) | Avg Latency (ms) | P95 Latency (ms) |
|-----|---------------|----------|----------------|-----|-------------|-----------------|
| 1 | Money Transfer | 1 shard (each shard 0.5 CPU, 1GB RAM) | 1m | **1445** | 69.15 | 94.29 |  
| 2 | Money Transfer | 2 shard (each shard 0.5 CPU, 1GB RAM) | 1m | **2383** | 41.92 | 78.03 |  
| 3 | Money Transfer | 2 shard (each shard 5.0 CPU, 6GB RAM) | 1m | **3987** | 25.06 | 39.26 |  
| 4 | pgbench | 1 database node (10 CPU, 12GB RAM) | 1m | **5565** | 17.96 | - |  

# 2. Observations & Analysis
## Money Transfer
- **Khi sử dụng 1 shard (0.5 CPU, 1GB RAM):**
   - Đạt 1445 TPS, độ trễ trung bình 69.15 ms, P95 latency 94.29 ms.

- **Khi mở rộng lên 2 shards (mỗi shard 0.5 CPU, 1GB RAM):**
   - TPS tăng lên 2383 TPS (tăng ~65% so với 1 shard).
   - Độ trễ trung bình giảm còn 41.92 ms, P95 latency giảm còn 78.03 ms.

- **Khi thêm resource cho 2 shards (mỗi shard 5 CPU, 6GB RAM):**
   - TPS đạt 3987 TPS, tăng thêm nhưng không tuyến tính với tài nguyên tăng.
   - Độ trễ trung bình giảm còn 25.06 ms, P95 latency giảm còn 39.26 ms.

## pgbench
- Với 1 database node (10 CPU, 12GB RAM), đạt 5565 TPS, độ trễ trung bình 17.96ms.
- Transaction đơn giản hơn, mô phỏng giao dịch ngân hàng, bao gồm: 1 SELECT, 3 UPDATE, 1 INSERT, không có thêm xử lý ngoài DB.

```sql
1. BEGIN;
2. UPDATE pgbench_accounts SET abalance = abalance + :delta WHERE aid = :aid;
3. SELECT abalance FROM pgbench_accounts WHERE aid = :aid;
4. UPDATE pgbench_tellers SET tbalance = tbalance + :delta WHERE tid = :tid;
5. UPDATE pgbench_branches SET bbalance = bbalance + :delta WHERE bid = :bid;
6. INSERT INTO pgbench_history (tid, bid, aid, delta, mtime) VALUES (:tid, :bid, :aid, :delta, CURRENT_TIMESTAMP);
7. COMMIT;
```

# 3. Detailed Analysis (oservation from monitoring chart)
My Test Machine: Macbook Pro M1 (10 CPU, 12GB RAM)

## Money Transfer vs. pgbench
- Money Transfer có TPS thấp hơn so với pgbench (cấu hình tương đương):
   - Ngữ cảnh Money Transfer:
      - Gồm 2 transaction con: mỗi transaction có 1 UPDATE, 1 INSERT.
      - Có thêm thao tác ghi Kafka (gây thêm I/O disk).
   - pgbench đơn giản hơn, không có xử lý ngoài DB.

=> Money Transfer chỉ đạt ~70% TPS so với pgbench, chủ yếu do có thêm thao tác ghi ngoài (Kafka) và nhiều transaction nhỏ.

## CPU Usage
Biểu đồ CPU không đạt mức bão hòa (<40% CPU test machine), chứng tỏ CPU không phải là nút thắt cổ chai chính (total 100% CPU).

Khi tăng số lượng CPU từ 0.5 → 10 CPU, mức sử dụng CPU tăng lên nhưng không tương ứng với mức tăng của TPS.

## Memory Usage
RAM sử dụng thấp (~4GB RAM), không có dấu hiệu bị thiếu bộ nhớ (total 12GB RAM).

Điều này gợi ý rằng vấn đề không phải do bộ nhớ bị hạn chế.

## Disk I/O và Network Traffic
- **I/O Metrics:**
   - **Context switch:** ~160k (rất cao)
   - **I/O Utilization:** ~80%
   - **Disk IO:** ~5k ops/s
   - **Time Spent on I/Os:** ~80%
   - **Disk R/W merged:** ~1.2k ops/s

Transaction có nhiều insert/update nhưng mỗi lần chỉ ghi một lượng nhỏ dữ liệu thay vì batch insert/update, việc ghi lên đĩa liên tục với cường độ cao gây ra disk I/O bottleneck.

## Conclusion
- **Nút thắt chính:**

   - **Disk I/O bottleneck:** Transaction nhỏ lẻ, liên tục ghi disk + Kafka log → Disk utilization cao (~80%), gây giới hạn TPS.

   - **Context switch quá nhiều (~160k):** Do xử lý nhiều transaction nhỏ, thread lifecycle ngắn, dẫn đến overhead scheduling.

   - **CPU và RAM không phải giới hạn chính**, dù CPU tăng mạnh nhưng TPS không tỷ lệ thuận, chủ yếu vì bottleneck nằm ở disk và context switching.

# Appendix
![Architecture Diagram](./performance-testing.png)

## Test 1. Testing transfer money on one shard (a single shard 0.5 CPU, 1GB RAM)
```sh
$ k6 run --vus 100 --duration 1m k6_scripts/transfer_money.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: k6_scripts/transfer_money.js
        output: -

     scenarios: (100.00%) 1 scenario, 100 max VUs, 1m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 1m0s (gracefulStop: 30s)

INFO[0060] Preparing the end-of-test summary...          source=console
INFO[0060]      ✓ is status 200

     checks.........................: 100.00% ✓ 86776       ✗ 0    
     data_received..................: 105 MB  1.7 MB/s
     data_sent......................: 58 MB   964 kB/s
     http_req_blocked...............: avg=5.2µs   min=0s     med=1µs     max=4.55ms  p(90)=3µs     p(95)=3µs    
     http_req_connecting............: avg=2.61µs  min=0s     med=0s      max=3.08ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=69.03ms min=1.55ms med=83.92ms max=1.04s   p(90)=91.73ms p(95)=94.17ms
       { expected_response:true }...: avg=69.03ms min=1.55ms med=83.92ms max=1.04s   p(90)=91.73ms p(95)=94.17ms
     http_req_failed................: 0.00%   ✓ 0           ✗ 86776
     http_req_receiving.............: avg=30.32µs min=5µs    med=20µs    max=16.9ms  p(90)=37µs    p(95)=52µs   
     http_req_sending...............: avg=15.69µs min=3µs    med=9µs     max=20.94ms p(90)=17µs    p(95)=22µs   
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=68.99ms min=1.52ms med=83.88ms max=1.04s   p(90)=91.68ms p(95)=94.12ms
     http_reqs......................: 86776   1445.402461/s
     iteration_duration.............: avg=69.15ms min=1.62ms med=84.02ms max=1.04s   p(90)=91.85ms p(95)=94.29ms
     iterations.....................: 86776   1445.402461/s
     vus............................: 100     min=100       max=100
     vus_max........................: 100     min=100       max=100  source=console
ERRO[0060] failed to handle the end-of-test summary      error="Could not save some summary information:\n\t- could not open 'json_results/transfer_money_VU_100_DURATION_1M0S.json': open json_results/transfer_money_VU_100_DURATION_1M0S.json: no such file or directory"

running (1m00.0s), 000/100 VUs, 86776 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m0s
```

## Test 2. Testing transfer money on two shard (each shard 0.5 CPU, 1GB RAM)
```sh
$ k6 run --vus 100 --duration 1m k6_scripts/transfer_money.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: k6_scripts/transfer_money.js
        output: -

     scenarios: (100.00%) 1 scenario, 100 max VUs, 1m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 1m0s (gracefulStop: 30s)

INFO[0060] Preparing the end-of-test summary...          source=console
INFO[0060]      ✓ is status 200

     checks.........................: 100.00% ✓ 143141      ✗ 0     
     data_received..................: 173 MB  2.9 MB/s
     data_sent......................: 96 MB   1.6 MB/s
     http_req_blocked...............: avg=3.76µs  min=0s     med=1µs     max=4.32ms   p(90)=2µs     p(95)=3µs    
     http_req_connecting............: avg=1.69µs  min=0s     med=0s      max=2.99ms   p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=41.81ms min=1.17ms med=38.12ms max=179.99ms p(90)=73.47ms p(95)=77.93ms
       { expected_response:true }...: avg=41.81ms min=1.17ms med=38.12ms max=179.99ms p(90)=73.47ms p(95)=77.93ms
     http_req_failed................: 0.00%   ✓ 0           ✗ 143141
     http_req_receiving.............: avg=26.51µs min=6µs    med=19µs    max=5.41ms   p(90)=33µs    p(95)=43µs   
     http_req_sending...............: avg=13.27µs min=3µs    med=9µs     max=7.98ms   p(90)=16µs    p(95)=19µs   
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=41.77ms min=1.14ms med=38.08ms max=179.95ms p(90)=73.43ms p(95)=77.89ms
     http_reqs......................: 143141  2383.512629/s
     iteration_duration.............: avg=41.92ms min=1.25ms med=38.23ms max=180.06ms p(90)=73.56ms p(95)=78.03ms
     iterations.....................: 143141  2383.512629/s
     vus............................: 100     min=100       max=100 
     vus_max........................: 100     min=100       max=100   source=console
ERRO[0060] failed to handle the end-of-test summary      error="Could not save some summary information:\n\t- could not open 'json_results/transfer_money_VU_100_DURATION_1M0S.json': open json_results/transfer_money_VU_100_DURATION_1M0S.json: no such file or directory"

running (1m00.1s), 000/100 VUs, 143141 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m0s
```

## Test 3. Testing transfer money on two shard (no limit resource [ 10 CPU, 12GB RAM ])
```sh
$ k6 run --vus 100 --duration 1m k6_scripts/transfer_money.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: testing/k6_scripts/transfer_money.js
        output: -

     scenarios: (100.00%) 1 scenario, 100 max VUs, 1m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 1m0s (gracefulStop: 30s)

INFO[0060] Preparing the end-of-test summary...          source=console
INFO[0060]      ✓ is status 200

     checks.........................: 100.00% ✓ 239393      ✗ 0     
     data_received..................: 283 MB  4.7 MB/s
     data_sent......................: 159 MB  2.7 MB/s
     http_req_blocked...............: avg=2.55µs  min=0s     med=1µs     max=4.18ms p(90)=2µs     p(95)=2µs    
     http_req_connecting............: avg=970ns   min=0s     med=0s      max=2.95ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=24.98ms min=1.18ms med=28.23ms max=1.05s  p(90)=36.12ms p(95)=39.19ms
       { expected_response:true }...: avg=24.98ms min=1.18ms med=28.23ms max=1.05s  p(90)=36.12ms p(95)=39.19ms
     http_req_failed................: 0.00%   ✓ 0           ✗ 239393
     http_req_receiving.............: avg=20.58µs min=5µs    med=17µs    max=4.58ms p(90)=29µs    p(95)=37µs   
     http_req_sending...............: avg=9.91µs  min=2µs    med=8µs     max=5.39ms p(90)=13µs    p(95)=17µs   
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s     p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=24.95ms min=1.15ms med=28.2ms  max=1.05s  p(90)=36.09ms p(95)=39.16ms
     http_reqs......................: 239393  3987.697809/s
     iteration_duration.............: avg=25.06ms min=1.24ms med=28.31ms max=1.05s  p(90)=36.19ms p(95)=39.26ms
     iterations.....................: 239393  3987.697809/s
     vus............................: 100     min=100       max=100 
     vus_max........................: 100     min=100       max=100   source=console
ERRO[0060] failed to handle the end-of-test summary      error="Could not save some summary information:\n\t- could not open 'json_results/transfer_money_VU_100_DURATION_1M0S.json': open json_results/transfer_money_VU_100_DURATION_1M0S.json: no such file or directory"

running (1m00.0s), 000/100 VUs, 239393 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m0s
```

## Test 4. Testing with pgbench
```sh
$ pgbench -i -s 10 mydb
dropping old tables...
creating tables...
generating data (client-side)...
vacuuming...                                                                                
creating primary keys...
done in 0.92 s (drop tables 0.03 s, create tables 0.02 s, client-side generate 0.63 s, vacuum 0.07 s, primary keys 0.17 s).
4eaa4b95da3d:~$ pgbench -c 100 -j 20 -T 60 mydb
pgbench (17.2)
starting vacuum...end.
transaction type: <builtin: TPC-B (sort of)>
scaling factor: 10
query mode: simple
number of clients: 100
number of threads: 20
maximum number of tries: 1
duration: 60 s
number of transactions actually processed: 334825
number of failed transactions: 0 (0.000%)
latency average = 17.968 ms
initial connection time = 35.468 ms
tps = 5565.384840 (without initial connection time)
```


