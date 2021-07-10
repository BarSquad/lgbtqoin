FROM golang:1.16

WORKDIR /app

COPY . /app
RUN go mod download
RUN pwd

RUN go build /app/server.go

CMD ["./server"]
EXPOSE 5300