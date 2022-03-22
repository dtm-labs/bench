English | [简体中文](./README-cn.md)
# XA AT comparison test
This test compares the performance of the dtm implementation of XA with the seata-golang implementation of AT.

## Test Environment
Host: MacBook Pro (15-inch, 2016)
CPU: 2.7 GHz quad-core Intel Core i7
Memory: 16 GB 2133 MHz LPDDR3
Storage: 512G SSD

Mysql: Docker installation of Mysql 5.7

## dtm Test Steps
#### Prepare Mysql
With docker-compose, start the following file in the dtm project

`https://github.com/dtm-labs/dtm/blob/main/helper/compose.store.yml`

After connecting to mysql, execute `bench.sql` in the current directory

#### Start dtm

`
git clone https://github.com/dtm-labs/dtm && cd dtm
go run main.go
`

#### Start the performance test service
`
git clone https://github.com/dtm-labs/bench.git
cd bench/xaat && go run main.go
`

#### Performance test
` ab -t 5 -c 3 localhost:8080/api/benchSuccess`

#### Results
`Requests per second: 9.42 [#/sec] (mean)`

#### Results without logs
Because logging to the console can take up a lot of performance, it is necessary to close or redirect the logs to the nul device in order to test performance. The results of the test after turning off the logging of dtm and bench are

`Requests per second: 16.42 [#/sec] (mean)`

## seata-golang test steps

Run the example as described in seata-go-samples (the version used for this test is `648ef0d`)

#### Execute the Performance test

`ab -t 5 -c 3 http://localhost:8003/createSoCommit`

#### Results
`Requests per second: 9.24 [#/sec] (mean)`

#### The result without log
After redirecting all relevant logs from seata-golang to the nul device, the result of the performance test is

`Requests per second: 13.38 [#/sec] (mean)`

## Performance analysis
See [AT mode](https://dtm.pub/practice/at) for details and performance analysis
