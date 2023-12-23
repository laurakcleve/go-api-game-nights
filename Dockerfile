FROM golang:1.22-rc-bookworm

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=1

RUN go build -o playedgamesapi

EXPOSE 4000

CMD ["./playedgamesapi"]