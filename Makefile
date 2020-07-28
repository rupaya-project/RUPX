.PHONY: rupaya rupaya-cross evm all test clean
.PHONY: rupaya-linux rupaya-linux-386 rupaya-linux-amd64 rupaya-linux-mips64 rupaya-linux-mips64le
.PHONY: rupaya-darwin rupaya-darwin-386 rupaya-darwin-amd64

GOBIN = $(shell pwd)/build/bin
GOFMT = gofmt
GO ?= 1.12
GO_PACKAGES = .
GO_FILES := $(shell find $(shell go list -f '{{.Dir}}' $(GO_PACKAGES)) -name \*.go)

GIT = git

rupaya:
	go run build/ci.go install ./cmd/rupaya
	@echo "Done building."
	@echo "Run \"$(GOBIN)/rupaya\" to launch rupaya."

gc:
	go run build/ci.go install ./cmd/gc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gc\" to launch gc."

bootnode:
	go run build/ci.go install ./cmd/bootnode
	@echo "Done building."
	@echo "Run \"$(GOBIN)/bootnode\" to launch a bootnode."

puppeth:
	go run build/ci.go install ./cmd/puppeth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/puppeth\" to launch puppeth."

all:
	go run build/ci.go install

test: all
	go run build/ci.go test

clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# Cross Compilation Targets (xgo)

rupaya-cross: rupaya-windows-amd64 rupaya-darwin-amd64 rupaya-linux
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-*

rupaya-linux: rupaya-linux-386 rupaya-linux-amd64 rupaya-linux-mips64 rupaya-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-*

rupaya-linux-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/rupaya
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-* | grep 386

rupaya-linux-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/rupaya
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-* | grep amd64

rupaya-linux-mips:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/rupaya
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-* | grep mips

rupaya-linux-mipsle:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/rupaya
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-* | grep mipsle

rupaya-linux-mips64:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/rupaya
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-* | grep mips64

rupaya-linux-mips64le:
	go run build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/rupaya
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-linux-* | grep mips64le

rupaya-darwin: rupaya-darwin-386 rupaya-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-darwin-*

rupaya-darwin-386:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/rupaya
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-darwin-* | grep 386

rupaya-darwin-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/rupaya
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-darwin-* | grep amd64

rupaya-windows-amd64:
	go run build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/rupaya
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/rupaya-windows-* | grep amd64
gofmt:
	$(GOFMT) -s -w $(GO_FILES)
	$(GIT) checkout vendor
