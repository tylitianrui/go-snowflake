# go-snowflake
twitter的snowflake算法 -- go实现
64位置，timestamp -毫秒级别

| sign | timestamp | datacenterID | machineID  |   sequence  | 
| -----| ----      | ----        |----         |----         |
|  1   |       41  |  5          |5            | 12

 将workerid划分为机房id+ 机器id   
 
 可根据需求划分，每部分所占用的bit位数
 
  ```text
    go  test  -bench=. -run=none   
    
    
    24405             46357 ns/
   ```
