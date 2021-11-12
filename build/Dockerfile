FROM golang:1.16.3-alpine3.13
WORKDIR /go/src/api

COPY go.mod go.sum ./

RUN go get -d -v ./...
RUN go install -v ./...

COPY . .

CMD cd $(mktemp -d); \
    go mod init tmp; \
    go get -v github.com/cespare/reflex; \
    cd /go/src/api; \
    reflex -r '(\.go$|go\.mod)' -s go run cmd/balance/main.go; 