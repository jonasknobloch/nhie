FROM golang:1.19 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go run ./cmd/build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/nhie

FROM alpine:3.17

COPY --from=builder /app/nhie /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 80
ENTRYPOINT ["/nhie"]