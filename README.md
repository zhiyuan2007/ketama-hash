# ketama-hash
implement ketama hash in golang

## Introduction

用go语言实现ketama思想的一致性哈希，测试服务器节点增加和减少时命中率能有多少。

## 主要问题
1. 如何生成指定数量的虚拟节点
```
对服务器s调用fnv哈希计算值hv、根据虚拟节点数n和无符号整数的最大值得出虚拟节点的步长step
根据hv + step X n 重新计算虚拟节点的hashv，生成key为hashv，值为s的map。这样虚拟节点到服务器的映射关系就构造完成。
```
2. 如何根据key快速查找对应的服务器节点
```
对map根据key排序，在排序后的数组中采用二分查找新key的hash值对应的位置，查找的条件是hash <= 数组值，
这样找到的位置，或者正好是虚拟节点，或者是顺时针走的时候遇到的第一个虚拟节点。
```
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
