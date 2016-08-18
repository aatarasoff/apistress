# apistress

This is very simple stress testing tool for API based on [vegeta](https://github.com/tsenart/vegeta)

## Motivation

Sometimes you want to check SLA of your API automatically. There are many tools for profiling or stressing you software but they have huge complexity or require analyse results by third party tools. This project helps you do it easy and in very simple way. One configuration file with very simple structure is needed. In configuration file you need to declare target, attack time, requests frequency and (this is more important) SLA for target service (99th latency percentile and percentage of successful requests). And if any test is failed, program print report and exit with error code. It gives very simple integration with different continuous delivery tools.

## Usage

Create file in current directory with name `config.json` in path where app is executed. File structure:
```
{
  "tests": [                                        //array of tests
    {
      "rps": 100,                                   //target request per second
      "duration": 1,                                //test duration
      "target": {
        "method": "POST",                           //http method for target url
        "url": "http://localhost:8080/test",        //target url
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
Then run compiled program.
