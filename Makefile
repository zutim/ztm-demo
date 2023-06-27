GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
LINUX=GOOS=linux GOARCH=amd64 $(GOBUILD)

build:
	#$(GOBUILD) -o sync_match -ldflags '-w -s' -race main.go

linuxbuild:
	$(LINUX) -o bin/app -ldflags '-w -s' cmd/main.go
	upx  bin/app


#protoc
generate:
	protoc --go_out=common/pb --go-grpc_out=common/pb common/pb/*.proto



