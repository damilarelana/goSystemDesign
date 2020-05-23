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
* [docker](https://www.docker.com)
* [iron/go](https://hub.docker.com/r/iron/go)
* `net/http`
* `fmt`
* `io`
* `log`
* `reflect`
* `errors`
* `os`
* `encoding/json`
* `strconv`

***

### Rerun
When using `go run` to run the code locally BUT without docker, remember to always delete `accounts.db` before re-run so as to ensure the DB is recreated each time - since `CreateBucket()` cannot handle a pre-existing bucket

*** 

### Go Binary
To build the go binary required by docker container, go to `/goSystemDesign/accountService` then run
```bash
    $ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o accountservice-linux-amd64
```

***

### Docker (via Go binary)

To create a docker image with the previously built binary, go to the project's root directory [`goSystemDesign`] and run:
```bash
    $ docker build -t systemdesign/accountservice-binary --file <full path to Dockerfile.binary> accountService/
```
where `<full path to Dockerfile.binary>` represents path to `Dockerfile.binary` e.g. `/home/go/src/github.com/someuser/goSystemDesign/accountService/Dockerfile.binary`. Then use created image `systemdesign/accountservice-binary` to initialize a docker container via command:
 ```bash
    $ docker run -d -p 8080:8080 systemdesign/accountservice-binary:latest
```
use `docker container ls` to validate existence of container with image `systemdesign/accountservice-binary:latest`. Finally test the service with url:
```bash
http://localhost:8080/accounts/10001
```
which should return some `json` data `{"id": "10001", "name": "Individual_1"}`.

***

### Docker (via raw Go code)
To create a docker image with raw go code, change directory to project's root `/goSystemDesign` and run:
```bash
    $ docker build -t systemdesign/accountservice-raw --file Dockerfile.raw .
```
then use created image `systemdesign/accountservice-raw` to initialize a docker container via command:
 ```bash
    $ docker run -d -p 8081:8080 systemdesign/accountservice-raw:latest
```
use `docker container ls` to validate existence of container with image `systemdesign/accountservice-raw:latest`. Finally test the service with the url
```bash
http://localhost:8081/accounts/10001
```
which should return some `json` data `{"id": "10001", "name": "Individual_1"}`.