### goSystemDesign

A simple Golang implementation of the system design of an Account microservice that consists of:

* Edge Services:
  - OAuth token relay
  - OAuth authorization server
  - configuration server
  - AMQP Messaging
  - Trace Analysis
  - Hystrix Stream Aggregation
  - Monitor Dashboard
* Core Service:
  - Security API
  - Account Composite Service
  - Image Service
  - VipService [verification and inspection]
  - Quotes Service

***

The code leverages the following packages:

* [boltDB](github.com/etcd-io/bbolt")
* [gorilla mux](github.com/gorilla/mux)
* [error](github.com/pkg/errors)
* [GoConvey](http://goconvey.co)
* `net/http`
* `fmt`
* `io`
* `log`
* `reflect`
* `errors`
* `os`
*	`encoding/json`
* `strconv`

***

### Rerun
remember to always delete `accounts.db` before re-run so as to ensure the DB is recreated each time - since `CreateBucket()` cannot handle a pre-existing bucket