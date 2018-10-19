# The dockerfile refers to the remote Calculator app
# The parent image the golang:alpine.
FROM golang:1.11-alpine

# The next step is to ADD the server and the calculator logic into the container
ADD server /go/src/awesomeProject/RemoteCalculator/server
ADD Calculator /go/src/awesomeProject/RemoteCalculator/Calculator
ADD remoteCalcpb /go/src/awesomeProject/RemoteCalculator/remoteCalcpb

# Following step is to RUN the path and built a go binary on the container
RUN apk update && apk add git
RUN go get google.golang.org/grpc

RUN go build awesomeProject/RemoteCalculator/server
RUN go install awesomeProject/RemoteCalculator/server

# The container doesn't have only one thing but for the moment I will follow the instructions on
# https://medium.com/@matzhouse/go-grpc-and-docker-c1fb2ec8e9f0
ENTRYPOINT ["/go/bin/server"]

# Finally I have to specify the port that the server is listening
EXPOSE 8024