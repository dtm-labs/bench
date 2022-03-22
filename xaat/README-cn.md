简体中文 | [English](./README.md)
# XA AT 对比测试
本测试主要用于对比dtm实现的XA和seata-golang实现的AT性能

## 测试环境
主机: MacBook Pro (15-inch, 2016)
CPU: 2.7 GHz 四核Intel Core i7
内存：16 GB 2133 MHz LPDDR3
存储：512G SSD

Mysql: Docker安装的Mysql5.7

## dtm 测试步骤
#### 准备Mysql
通过docker-compose，启动dtm项目里的下面文件

`https://github.com/dtm-labs/dtm/blob/main/helper/compose.store.yml`

连接上mysql后，执行当前目录下的`bench.sql`

#### 启动dtm

`
git clone https://github.com/dtm-labs/dtm && cd dtm
go run main.go
`

#### 启动压测服务
`
git clone https://github.com/dtm-labs/bench.git
cd bench/xaat && go run main.go
`

#### 压测
`ab -t 5 -c 3 localhost:8080/api/benchSuccess`

#### 结果
`Requests per second:    9.42 [#/sec] (mean)`

#### 关闭日志后压测
因为日志输出到控制台会占用大量的性能，因此为了实测性能，需要将日志关闭或者重定向到nul设备。关闭dtm和bench之后的压测结果为：

`Requests per second:    16.42 [#/sec] (mean)`

## seata-golang 测试步骤

按照seata-go-samples的说明，运行例子（本次测试使用的版本为`648ef0d`）

#### 执行压测

`ab -t 5 -c 3  http://localhost:8003/createSoCommit`

#### 结果
`Requests per second:    9.24 [#/sec] (mean)`

#### 关闭日志之后的结果为
将seata-golang所有相关的日志全部重定向到nul设备后，压测的结果为：

`Requests per second:    13.38 [#/sec] (mean)`

## 性能分析
详细原理以及性能分析参考 [AT模式](https://dtm.pub/practice/at)