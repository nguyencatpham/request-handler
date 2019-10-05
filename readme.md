# Request Handler

[![N|Solid](https://cldup.com/dTxpPi9lDf.thumb.png)](https://github.com/nguyencatpham/request-handler)

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://github.com/nguyencatpham/request-handler)

Request handler is a package extend from net/http. It provides HTTP client and server implementations easier.
  - Get, Head, Post, and PostForm make HTTP (or HTTPS) requests
  - Clients and Transports are safe for concurrent use by multiple goroutines and for efficiency should only be created once and re-used.

# Getting Started!

  #### Model input
  ```
  type RequestModel struct {
	URL       string
	TokenType TokenType
	Token     string
	Username  string
	Password  string
	Body      string
}
  ```

### Installation

Request handler requires [Go](https://golang.org/) v1.11+ to run.

Install the package.

```sh
$ go get github.com/nguyencatpham/request-handler
```

#### Kubernetes + Google Cloud

See [KUBERNETES.md](https://github.com/joemccann/dillinger/blob/master/KUBERNETES.md)


### Todos

 - Write MORE Tests
 - Add Night Mode

License
----

MIT

