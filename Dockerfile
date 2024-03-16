FROM golang:1.21.6

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN rm -rf .env

RUN go build -v -o bin/app ./cmd/main.go

CMD ["./bin/app"]
