PKG_OS = $$(echo ${GOOS:=darwin})
PKG_ARCH = $$(echo ${GOARCH:=amd64})

.PHONY: all
all: clean build benchmark

.PHONY: clean
clean:
	@go clean

.PHONY: build
build:
	@GOOS=$(PKG_OS) GOARCH=$(PKG_ARCH) go build -a -o pkg/damm-$(PKG_OS)-$(PKG_ARCH) src/damm.go src/main.go

.PHONY: test
test:
	@go test -v ./src/...

.PHONY: benchmark
benchmark:
	@go test -v ./src/... -bench=.
