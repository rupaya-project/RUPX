.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/rupaya-project/rupx/cmd/swarm
	go build -o bin/rupaya     github.com/rupaya-project/rupx/cmd/rupaya
	go build -o bin/bootnode github.com/rupaya-project/rupx/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/rupaya-project/rupx/swarm/...
