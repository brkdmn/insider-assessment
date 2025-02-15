FROM golang:1.23 AS builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/api/main.go -o docs

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/api cmd/api/main.go

FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/docs /root/docs
COPY --from=builder /go/bin/api .
RUN chmod +x api
CMD ["./api"] 