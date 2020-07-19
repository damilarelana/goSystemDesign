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
* [docker-machine](https://docs.docker.com/machine/install-machine/)
* [docker swarm](https://docs.docker.com/engine/swarm/)
* [docker swarm visualizer](https://github.com/dockersamples/docker-swarm-visualizer)
* [Dvizz](https://github.com/eriklupander/dvizz)
* [iron/go](https://hub.docker.com/r/iron/go)
* [vagrant](https://www.vagrantup.com/)
* [ansible](https://www.ansible.com/community)
* [kubernetes](https://kubernetes.io/)
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

### Docker (via Go binary) [recommended]

To create a docker image with the previously built binary, go to the project's root directory [`goSystemDesign`] and run:
```bash
$ docker build -t systemdesign/accountservice-binary --file <full path to Dockerfile.binary> accountService/
``` 
where `<full path to Dockerfile.binary>` represents path to `Dockerfile.binary` e.g. `/home/go/src/github.com/someuser/goSystemDesign/accountService/Dockerfile.binary`. Then use created image `systemdesign/accountservice-binary` to initialize a docker container via command:
```bash
$ docker run -d -p 8081:8080 systemdesign/accountservice-binary:latest
```
Use `docker container ls` to validate existence of container with image `systemdesign/accountservice-binary:latest`. Finally test the service with url:
```bash
$ http://localhost:8080/accounts/10001
```
which should return some `json` data `{"id": "10001", "name": "Individual_1"}`.

***

### Docker (via raw Go code) [not recommended]

To create a docker image with raw go code, change directory to project's root `/goSystemDesign` and run:
```bash
$ docker build -t systemdesign/accountservice-raw --file Dockerfile.raw .
```
Then use created image `systemdesign/accountservice-raw:latest` to initialize a docker container via command:
```bash
$ docker run -d -p=8081:8081 systemdesign/accountservice-raw:latest
```
The internal port `8080` (from the `8080:8081` was made to match the port exposed by the docker image `systemdesign/accountservice-raw:latest`). Use `docker container ls` to validate existence of container with image `systemdesign/accountservice-raw:latest`. Finally test the service with the url:
```bash
$ http://localhost:8081/accounts/10001
```
which should return some `json` data `{"id": "10001", "name": "Individual_1"}`.

***

### Docker Swarm

Initialize docker swarm mode: `docker swarm init --advertise-addr 192.168.15.2;2377`
Create docker overlay network: `docker network create --driver overlay systemdesign-network`
Deploy `accountservice` to the swarm: 
```
$ docker service create --replicas=1 --name=systemdesign-accountservice  --network=systemdesign-network -p=8081:8080 systemdesign/accountservice-binary:latest
```
Use `docker service ls` to validate that the service is running on the swarm. Get the ip-address of the manager node by running `docker node ls` to the node `ID` for the `Leader`. Finally test the service with the url:
```bash
$ http://127.0.0.1:8080/accounts/10001
```
which should return some `json` data `{"id": "10001", "name": "Individual_1"}`. Note: it is assumed that `docker swarm init` was used to create the new swarm (in the local machine), and not `docker swarm init --advertise-address <MANAGER-IP>`, consequently the service is externally available at `127.0.0.1`

***

### Swarm Visualization

Visualize existing services on docker swarm, by creating an additional service using [docker swarm visualizer](https://github.com/dockersamples/docker-swarm-visualizer) ):
```
$ docker service create --name=viz --replicas=1 --publish=8080:8000/tcp --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock --network=systemdesign-network dockersamples/visualizer
```
check http://127.0.0.1:8000 for the swarm visualization. Alternatively, [dvizz](https://github.com/eriklupander/dvizz) can also be instead of `docker swarm visualizer`:
```
$ docker service create --name=dvizz --replicas=1 --publish=6969:8001 --constraint=node.role==manager --mount=type=bind,src=/var/run/docker.sock,dst=/var/run/docker.sock --network=systemdesign-network dvizz:latest
```
check http://127.0.0.1:8001 for the alternative swarm visualization.

***

### Optional

To further help the visualization content and philosophy of microservice, a `java` based service (in [docker hub](https://hub.docker.com/r/eriklupander/quotes-service/)) can be deployed as `quoteservice` to the swarm:
```
$ docker service create --replicas=1 --name=systemdesign-quoteservice  --network=systemdesign-network eriklupander/quotes-service
```