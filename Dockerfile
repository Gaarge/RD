FROM golang:1.20-alpine

WORKDIR /app

COPY go ./go
COPY templates ./templates
COPY static ./static

WORKDIR /app/go

RUN go mod init site

EXPOSE 8080

CMD ["go", "run", "main.go"]
