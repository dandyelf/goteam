# go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
GOPATH := $(shell go env GOPATH)

all: prot serv

prot:
	export PATH="$(PATH):$(GOPATH)/bin" && \
	cd prot && \
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative my.proto

testl:
	cd client && go run .

serv:
	cd server && go run .


#         protoc -I./pkg/server \
                --go_out ./pkg/server \
                --go_opt=paths=source_relative \
                --go-grpc_out ./pkg/server \
                --go-grpc_opt=paths=source_relative \
                ./pkg/server/*.proto