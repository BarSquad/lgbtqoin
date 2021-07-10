FROM golang:1.16

WORKDIR /app

COPY . /app

RUN go build server.go

CMD ["./server"]
EXPOSE 5300