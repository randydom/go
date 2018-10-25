EXE=main
IMAGE=filestream-server

all:
	make dependencies
	make local
	make docker

dependencies:
	go get github.com/golang/protobuf/proto
	go get golang.org/x/net/context
	go get google.golang.org/grpc
	go get golang.org/x/sys/unix

protobuf:
	/usr/local/bin/protoc -I ./ fileShare.proto --go_out=plugins=grpc:./

local:
	go build -o $(EXE) main.go

docker:
	make docker-image
	make docker-local

static:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $(EXE) main.go

docker-image:
	make static
	docker build -t $(IMAGE) .

docker-local:
	docker run -d -p 8080:8080/tcp -p 9000:9000/tcp $(IMAGE)

