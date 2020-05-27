FROM golang:1.14 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/nhie

FROM scratch
COPY --from=builder /app/nhie /
EXPOSE 8080
ENTRYPOINT ["/nhie"]