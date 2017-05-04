# ketama-hash
implement ketama hash in golang

## Introduction

用go语言实现ketama思想的一致性哈希，测试服务器节点增加和减少时命中率能有多少。

## Synopsis

go run ketama_hash.go
    
## result

```
begin server  3.3.3.3  has key count  295
begin server  1.1.1.1  has key count  374
begin server  2.2.2.2  has key count  331
add server  2.2.2.2  has key count  241
add server  1.1.1.1  has key count  269
add server  4.4.4.4  has key count  250
add server  3.3.3.3  has key count  240
reduce server  1.1.1.1  has key count  509
reduce server  2.2.2.2  has key count  491
-----------add a new server------------hit ratio--------------------
server 1.1.1.1, hit 269, total 374, hit_ratio 0.72
server 2.2.2.2, hit 241, total 331, hit_ratio 0.73
server 3.3.3.3, hit 240, total 295, hit_ratio 0.81
-----------reduce a server------------hit ratio--------------------
server 2.2.2.2, hit 331, total 491, hit_ratio 0.67
server 1.1.1.1, hit 374, total 509, hit_ratio 0.73
```
## Author

ben <liuben5918@gmail.com>
