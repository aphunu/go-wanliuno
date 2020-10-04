# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: wanliuno android ios wanliuno-cross evm all test clean
.PHONY: wanliuno-linux wanliuno-linux-386 wanliuno-linux-amd64 wanliuno-linux-mips64 wanliuno-linux-mips64le
.PHONY: wanliuno-linux-arm wanliuno-linux-arm-5 wanliuno-linux-arm-6 wanliuno-linux-arm-7 wanliuno-linux-arm64
.PHONY: wanliuno-darwin wanliuno-darwin-386 wanliuno-darwin-amd64
.PHONY: wanliuno-windows wanliuno-windows-386 wanliuno-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

wanliuno:
	$(GORUN) build/ci.go install ./cmd/wanliuno
	@echo "Done building."
	@echo "Run \"$(GOBIN)/wanliuno\" to launch wanliuno."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/wanliuno.aar\" to use the library."

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go get -u golang.org/x/tools/cmd/stringer
	env GOBIN= go get -u github.com/kevinburke/go-bindata/go-bindata
	env GOBIN= go get -u github.com/fjl/gencodec
	env GOBIN= go get -u github.com/golang/protobuf/protoc-gen-go
	env GOBIN= go install ./cmd/abigen
	@type "npm" 2> /dev/null || echo 'Please install node.js and npm'
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

wanliuno-cross: wanliuno-linux wanliuno-darwin wanliuno-windows wanliuno-android wanliuno-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-*

wanliuno-linux: wanliuno-linux-386 wanliuno-linux-amd64 wanliuno-linux-arm wanliuno-linux-mips64 wanliuno-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-*

wanliuno-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/wanliuno
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep 386

wanliuno-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/wanliuno
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep amd64

wanliuno-linux-arm: wanliuno-linux-arm-5 wanliuno-linux-arm-6 wanliuno-linux-arm-7 wanliuno-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep arm

wanliuno-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/wanliuno
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep arm-5

wanliuno-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/wanliuno
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep arm-6

wanliuno-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/wanliuno
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep arm-7

wanliuno-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/wanliuno
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep arm64

wanliuno-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/wanliuno
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep mips

wanliuno-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/wanliuno
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep mipsle

wanliuno-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/wanliuno
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep mips64

wanliuno-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/wanliuno
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-linux-* | grep mips64le

wanliuno-darwin: wanliuno-darwin-386 wanliuno-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-darwin-*

wanliuno-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/wanliuno
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-darwin-* | grep 386

wanliuno-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/wanliuno
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-darwin-* | grep amd64

wanliuno-windows: wanliuno-windows-386 wanliuno-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-windows-*

wanliuno-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/wanliuno
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-windows-* | grep 386

wanliuno-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/wanliuno
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/wanliuno-windows-* | grep amd64
