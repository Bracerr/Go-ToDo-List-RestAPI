FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o todo-app ./src/cmd/main.go

FROM alpine:3.17

RUN apk update && apk --no-cache add ca-certificates

COPY --from=builder /app/todo-app /usr/local/bin/todo-app

RUN chmod +x /usr/local/bin/todo-app

COPY .env ./

ENV PORT=8080

EXPOSE 8080

CMD ["todo-app"]
