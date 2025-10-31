FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o api ./cmd/api

EXPOSE ${HTTP_ADDR}

CMD ["./api"]