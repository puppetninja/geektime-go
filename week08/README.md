# Q&A

1、使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

10 bytes
```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 10 -q 
SET: 113765.64 requests per second, p50=0.215 msec                    
GET: 110497.24 requests per second, p50=0.231 msec

Summary:
  throughput summary: 106837.61 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.239     0.072     0.239     0.287     0.391     0.879
Summary:
  throughput summary: 105932.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.241     0.104     0.231     0.351     0.431     0.903
```

20 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 20 -q
SET: 106609.80 requests per second, p50=0.231 msec                    
GET: 101626.02 requests per second, p50=0.239 msec


Summary:
  throughput summary: 101936.80 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.251     0.112     0.239     0.367     0.431     0.815
Summary:
  throughput summary: 107411.38 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.237     0.112     0.231     0.303     0.399     0.911
```

50 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 50 -q
SET: 105485.23 requests per second, p50=0.231 msec                    
GET: 109409.20 requests per second, p50=0.231 msec

Summary:
  throughput summary: 103412.62 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.247     0.096     0.239     0.335     0.423     0.999

Summary:
  throughput summary: 104493.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.244     0.096     0.239     0.343     0.439     0.855

```

100 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 100 -q
SET: 110375.27 requests per second, p50=0.231 msec                    
GET: 107526.88 requests per second, p50=0.231 msec

Summary:
  throughput summary: 105042.02 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.243     0.096     0.239     0.327     0.423     0.807

Summary:
  throughput summary: 101832.99 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.250     0.088     0.239     0.335     0.439     0.807
```

200 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 200 -q
SET: 109649.12 requests per second, p50=0.231 msec                    
GET: 110497.24 requests per second, p50=0.223 msec

Summary:
  throughput summary: 106723.59 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.239     0.096     0.231     0.303     0.415     0.863

Summary:
  throughput summary: 111856.82 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.229     0.120     0.223     0.311     0.423     0.879
```

1k bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 1000 -q
SET: 105042.02 requests per second, p50=0.239 msec                    
GET: 111358.58 requests per second, p50=0.215 msec


Summary:
  throughput summary: 100401.61 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.255     0.072     0.239     0.343     0.527     4.047

Summary:
  throughput summary: 112994.35 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.227     0.064     0.215     0.311     0.431     1.159

```

5k bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 5000 -q
SET: 101832.99 requests per second, p50=0.247 msec                   
GET: 111358.58 requests per second, p50=0.223 msec

Summary:
  throughput summary: 98619.32 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.259     0.128     0.255     0.327     0.415     1.135

Summary:
  throughput summary: 104275.29 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.246     0.104     0.239     0.335     0.447     0.831
```

随着value增大，从平均分布上（p95）来讲，性能损失几乎没有, p100个别在value较大1k有较大延迟。

2、写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

10 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 10 -q -n 100000
SET: 106951.88 requests per second, p50=0.239 msec                    
GET: 108932.46 requests per second, p50=0.231 msec  


23) "keys.bytes-per-key"
24) (integer) 62128
```

20 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 20 -q -n 100000
SET: 107642.62 requests per second, p50=0.231 msec                    
GET: 103734.44 requests per second, p50=0.239 msec

23) "keys.bytes-per-key"
24) (integer) 62136
```

50 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 50 -q -n 100000
SET: 108459.87 requests per second, p50=0.231 msec                    
GET: 105820.11 requests per second, p50=0.239 msec     

23) "keys.bytes-per-key"
24) (integer) 62168
```

100 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 100 -q -n 100000
SET: 106496.27 requests per second, p50=0.239 msec                    
GET: 106723.59 requests per second, p50=0.239 msec 

23) "keys.bytes-per-key"
24) (integer) 62224
```

200 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 200 -q -n 100000
SET: 102774.92 requests per second, p50=0.231 msec                   
GET: 113378.68 requests per second, p50=0.223 msec   

23) "keys.bytes-per-key"
24) (integer) 3158144
```

1000 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 1000 -q -n 100000
SET: 105152.48 requests per second, p50=0.239 msec                    
GET: 116550.12 requests per second, p50=0.215 msec

23) "keys.bytes-per-key"
24) (integer) 63136
```

5000 bytes

```sh
root@a7fda070cab7:/data# redis-benchmark -t set,get -d 5000 -q -n 100000
SET: 98716.68 requests per second, p50=0.255 msec                    
GET: 106609.80 requests per second, p50=0.231 msec 

23) "keys.bytes-per-key"
24) (integer) 67232
```
