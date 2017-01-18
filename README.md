# apistress

[![Go Report Card](https://goreportcard.com/badge/github.com/aatarasoff/apistress)](https://goreportcard.com/report/github.com/aatarasoff/apistress)

This is very simple stress testing tool for API based on [vegeta](https://github.com/tsenart/vegeta)

## Motivation

Sometimes you want to check SLA of your API automatically. There are many tools for profiling or stressing your software but all of them have huge complexity or require analyse results by third party tools. This project helps you do it easy and in very simple way. One configuration file with very simple structure is needed. In configuration file you need to declare target, attack time, requests frequency and (this is most important part) SLA for target service (99th latency percentile and percentage of successful requests). And if any test is failed, program will print report and exit with error code. It gives very simple integration with different continuous delivery tools.

## Usage

Create file with name `config.json` and folowing structure:
```Go
{
  "baseUrl": "http://localhost:8080",               //base url for targets
  "tests": [                                        //array of tests
    {
      "rps": 100,                                   //target request per second
      "duration": 1,                                //test duration
      "target": {
        "method": "POST",                           //http method for target url
        "path": "/test",                            //relative path
        "headers": [                                //optional
          {
            "name": "X-Request-Id",                 //http header name
            "value": "12345"                        //http header value
          }
        ],
        "body": "ewoic2F5IiA6ICJoZWxsbyIKfQ=="      //base64 endoded request body (optional)
      },
      "sla": {
        "latency": 150,                             //99 percenitle latency
        "successRate": 99.9                         //percentage of successful requests (2xx http code is returned)
      }
    }
  ]
}
```
Then run docker container:
```bash
docker run --rm --net=host \
   -v /path/to/folder/with/config:/data \
   aatarasoff/apistress
```
or with overriden `baseUrl` config property:
```bash
docker run --rm --net=host \
   -v /path/to/folder/with/config:/data \
   aatarasoff/apistress apistress \
   -baseUrl http://custom.server:8080
```
If `stdin` input is required, use `-config=stdin` flag:
```bash
cat config.json | docker run --rm --net=host \
   aatarasoff/apistress apistress \
   -config=stdin
```
Also it is possible to define own config file name and path:
```bash
docker run --rm --net=host \
   -v /path/to/folder/with/config:/data \
   aatarasoff/apistress apistress \
   -config=/path/to/folder/with/config/filename.json
```
For each test program prints metrics into `stdout`:
```bash
Requests      [total, rate]            10, 11.11
Duration      [total, attack, wait]    1.007898226s, 899.953255ms, 107.944971ms
Latencies     [mean, 50, 95, 99, max]  108.246432ms, 107.608893ms, 109.534083ms, 109.534083ms, 112.276495ms
Bytes In      [total, mean]            5750, 575.00
Bytes Out     [total, mean]            0, 0.00
Success       [ratio]                  100.00%
Status Codes  [code:count]             200:10
Error Set:
```
and returns exit code that you may check with following command:
```bash
echo $?       //0 - ok, 1 - sla error
```

## Usage without Docker or developing

You need install and setup `golang` 1.6 or above with following [instructions](https://golang.org/doc/install). Then run:
```Go
go get
go install github.com/aatarasoff/apistress
```
