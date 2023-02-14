SRCDIR := api
DSTDIR := pkg/api

.PHONY: all
all: pkg/api/user.pb.go

pkg/api/user.pb.go: api/user.proto
	protoc -I ./api --go_out=pkg/api ./api/user.proto