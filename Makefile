.PHONY: $(BIN) test doc

export GOPATH:=$(shell pwd)
export CWD:=$(shell pwd)

# инжектим в переменную version_info данные о сборке
#LDFLAGS="-X main.version_info \"$(shell git rev-parse --short HEAD) built at $(shell date) on $(shell hostname)\""

appleworker:
	go install appleworker

client:
	go install client

run_appleworker: appleworker
	bin/appleworker

run_client: client
	bin/client

test:
	go test $(BIN) -config $(CWD)/etc/unittest.json

doc:
	godoc -http=:6060 -goroot=`pwd`

jenkins:
	go test $(BIN) -v -config $(CWD)/etc/unittest.json

fmt:
	go fmt $(BIN)

depends:
	go get github.com/anachronistic/apns
	go get github.com/streadway/amqp
