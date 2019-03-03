CRUDWEBSOCKET_PKG_VERSION?=0.0.0
COMMIT=`git rev-parse --short HEAD`

build: 
	go install -v --ldflags "-w \
	-X github.com/mrcaelumn/crudwebsocket/version.Version/version.Version=$(CRUDWEBSOCKET_PKG_VERSION) \
	-X github.com/mrcaelumn/crudwebsocket/version.Version/version.GitCommit=$(COMMIT)" .

build_binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o crudwebsocket -a --ldflags "-w \
	-X github.com/mrcaelumn/crudwebsocket/version.Version=$(CRUDWEBSOCKET_PKG_VERSION) \
	-X github.com/mrcaelumn/crudwebsocket/version.GitCommit=$(COMMIT)" .

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet -v $(shell go list ./... | grep -v /vendor/)

clean:
	@rm -rf build
	@rm -rf crudwebsocket*

.PHONY: test vet build build_binary clean
