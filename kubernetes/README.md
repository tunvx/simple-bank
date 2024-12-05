# Resources Allocated To Services

**Table Estimating Resources To Be Allocated:**

| Pods/Deployment | Replicas | CPU (cores) | MEM (Gi) | Disk (Gi) |
|-----------------|--------------|-----|-----|-----|
| postgres01 (core_db) | 1  | 4 | 8 | 50 | 
| postgres02 (auth_db) | 1  | 1 | 1 | 50 | 
| redis (cache, mq)    | 1  | 1 | 1 | 10 | 
| manage_service       | 1  | 0.5 | 0.25 | 5 |  
| auth_service         | 1  | 0.5 | 0.25 | 5 |  
| notification_service | 1  | 0.5 | 0.25 | 5 | 
| transaction_service  | 5* | 0.5 | 0.25 | 5 |  
|----------------------|---|-----|------|-----|
| total_used           | * | 10  | 12   | 150 | 
| resouce_all          | * | 10  | 13.2 | 256 | 