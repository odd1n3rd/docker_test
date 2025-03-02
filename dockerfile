FROM golang:1.24-alpine

WORKDIR /app

COPY main.go .

RUN go build -o docker-go-hello main.go

CMD [ "./docker-go-hello" ]