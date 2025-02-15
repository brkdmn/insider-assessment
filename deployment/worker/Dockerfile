FROM golang:1.23 AS builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/worker cmd/worker/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/configs /app/configs

COPY --from=builder /go/bin/worker .
RUN chmod +x worker
CMD ["./worker"] 