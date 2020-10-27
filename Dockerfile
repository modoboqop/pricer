FROM golang:1.15

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPATH=/

COPY go.mod .
COPY go.sum .

COPY . .

RUN go build -o main .

CMD ["./main"]

