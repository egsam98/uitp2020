FROM golang:alpine

RUN apk update && \
    apk add build-base && \
    apk add --no-cache git

WORKDIR /github.com/egsam98/uitp

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o bin/uitp *.go
EXPOSE 8080
ENTRYPOINT ["bin/uitp"]