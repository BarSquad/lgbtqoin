FROM golang:1.16

WORKDIR /app

COPY . /app
RUN pwd
RUN ls
RUN ls /app/proto
RUN go get -u google.golang.org/grpc@latest
RUN go get -u google.golang.org/grpc/grpclog@latest
RUN go build ./server.go

CMD ["./server"]
EXPOSE 5300