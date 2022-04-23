FROM golang:1.16-alpine

WORKDIR ./ app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg

CMD ["go", "run", "./cmd/main.go"]
