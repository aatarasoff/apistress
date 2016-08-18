# apistress

This is very simple stress testing tool for API

## Motivation

//TODO

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
