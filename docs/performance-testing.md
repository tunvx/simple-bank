# 1. Testing transfer money on one shard (a single shard 0.5 CPU, 1GB RAM)
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

# 2. Testing transfer money on two shard (each shard 0.5 CPU, 1GB RAM)
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

# 3. Testing transfer money on two shard (no limit resource [ 10 CPU, 12GB RAM ])
```sh
$ k6 run --vus 100 --duration 2m k6_scripts/transfer_money.js


          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

     execution: local
        script: k6_scripts/transfer_money.js
        output: -

     scenarios: (100.00%) 1 scenario, 100 max VUs, 2m30s max duration (incl. graceful stop):
              * default: 100 looping VUs for 2m0s (gracefulStop: 30s)

INFO[0120] Preparing the end-of-test summary...          source=console
INFO[0120]      ✓ is status 200

     checks.........................: 100.00% ✓ 445727      ✗ 0     
     data_received..................: 526 MB  4.4 MB/s
     data_sent......................: 297 MB  2.5 MB/s
     http_req_blocked...............: avg=2.21µs  min=0s     med=1µs     max=15.22ms  p(90)=2µs     p(95)=3µs    
     http_req_connecting............: avg=512ns   min=0s     med=0s      max=3.1ms    p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=26.82ms min=1.87ms med=29.37ms max=564.52ms p(90)=36.85ms p(95)=39.61ms
       { expected_response:true }...: avg=26.82ms min=1.87ms med=29.37ms max=564.52ms p(90)=36.85ms p(95)=39.61ms
     http_req_failed................: 0.00%   ✓ 0           ✗ 445727
     http_req_receiving.............: avg=23.52µs min=5µs    med=19µs    max=14.51ms  p(90)=30µs    p(95)=36µs   
     http_req_sending...............: avg=12.02µs min=3µs    med=8µs     max=20.64ms  p(90)=15µs    p(95)=17µs   
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=26.78ms min=1.84ms med=29.33ms max=560.75ms p(90)=36.81ms p(95)=39.56ms
     http_reqs......................: 445727  3713.634673/s
     iteration_duration.............: avg=26.91ms min=1.95ms med=29.46ms max=564.72ms p(90)=36.94ms p(95)=39.7ms 
     iterations.....................: 445727  3713.634673/s
     vus............................: 100     min=100       max=100 
     vus_max........................: 100     min=100       max=100   source=console

running (2m00.0s), 000/100 VUs, 445727 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  2m0s
```

# 4. Testing with pgbench
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

# 5. Monitoring four test
![Architecture Diagram](./performance-testing-monitoring.png)

# 6. Phân tích kết quả
My Test Machine: Macbook Pro M1 (10 CPU, 12GB RAM)

### 1. Kết quả kiểm thử
- **1 shard (0.5 CPU, 1GB RAM):** 1445 TPS, avg latency 69.15ms.
- **2 shards (0.5 CPU, 2GB RAM each):** 2383 TPS, avg latency 41.92ms.
- **2 shards (10 CPU, 12GB RAM each):** 3713 TPS, avg latency 26.91ms.
    - **Context 1:** 1 transaction: 2 update, 2 insert.
    - **Context 2:**
        - **1 kafka write:** 1 insert
        - **1 transaction:** 1 update, 1 insert.
        - **1 transaction:** 1 update, 1 insert.
- pgbench test (10 CPU, 12GB RAM each): 5565 TPS, avg latency 17.968ms.
    - 1 transaction: 1 SELECT, 3 UPDATE, 1 INSERT

Nhìn qua, việc tăng số lượng CPU và RAM giúp giảm độ trễ trung bình (latency), nhưng TPS không tăng đáng kể khi tài nguyên tăng lên nhiều lần.
- 5565 TPS (pgbench) * 2/3 = ~3713 TPS (transfer money) (vì ngữ cảnh chuyển tiền có thêm 1 lần kafka ghi xuống đĩa nữa)

### 2. Quan sát từ biểu đồ monitoring
**a) CPU Usage**
Biểu đồ CPU không đạt mức bão hòa (có phần trống), chứng tỏ CPU không phải là nút thắt cổ chai chính.
Khi tăng số lượng CPU từ 0.5 → 10, mức sử dụng CPU tăng lên nhưng không tương ứng với mức tăng của TPS.
**b) Memory Usage**
RAM sử dụng thấp, không có dấu hiệu bị thiếu bộ nhớ.
Điều này gợi ý rằng vấn đề không phải do bộ nhớ bị hạn chế.
**c) Disk I/O và Network Traffic**
Nếu có nhiều update nhưng mỗi lần chỉ ghi một lượng nhỏ dữ liệu, việc ghi đĩa liên tục có thể gây ra context switch cao (~160k) và gây tắc nghẽn I/O (IO utilization:: 80%, Disk IO:: 5k io/s).
Nếu transaction ghi xuống DB theo từng bản ghi nhỏ thay vì batch insert/update, số lần gọi I/O sẽ quá nhiều.
### 3. Kết luận
Nút thắt chính có thể là context switch quá nhiều. Khi cập nhật từng record nhỏ lẻ, số lượng syscalls và context switch tăng mạnh, khiến hiệu suất bị giới hạn.
Disk I/O có thể là vấn đề. Nếu DB liên tục ghi dữ liệu theo từng giao dịch nhỏ mà không batch lại, tốc độ ghi của ổ đĩa sẽ giới hạn TPS.
CPU không phải vấn đề lớn, nhưng overhead scheduling có thể là nguyên nhân. Việc tăng CPU từ 0.5 lên 10 có thể dẫn đến nhiều thread hơn, nhưng nếu mỗi thread chỉ làm việc rất ngắn rồi chuyển context, hệ thống sẽ mất thời gian để quản lý hơn là thực sự xử lý logic.

### 4. AWS Limitatioin
- **pgbench:** [41,498 transactions with Amazon Aurora PostgreSQL-compatibleedition dr4.16xlarge - 64 vCPUs, 488 GiB.](https://d1.awsstatic.com/product-marketing/Aurora/RDS_Aurora_PostgreSQL_Performance_Assessment_Benchmarking_V1-0.pdf)


