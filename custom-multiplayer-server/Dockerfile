FROM golang:1.23.2

WORKDIR /app

COPY . .

RUN go build -o multiplayer-server ./cmd/app

EXPOSE 8080/udp
EXPOSE 8081/tcp

CMD ["/app/multiplayer-server"]