FROM golang:1.16-alpine

WORKDIR /app

COPY . .

# RUN go mod download
RUN go get ./...
# RUN apk update && apk add bash && bash setup-env.sh docker
# RUN go build -o cmd/main.go

CMD ["go", "run", "cmd/main.go"]
