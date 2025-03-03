FROM golang:1.24-alpine AS build

WORKDIR /app

COPY main.go .

RUN go build -o docker-go-hello main.go


FROM alpine:3.15

WORKDIR /app

COPY --from=build  /app/docker-go-hello .

CMD [ "./docker-go-hello" ]