FROM golang:1.21

RUN mkdir /app
WORKDIR /app

COPY server.go /app/server.go

RUN go build -o /app/server /app/server.go

EXPOSE 8080

CMD ["/app/server"]
