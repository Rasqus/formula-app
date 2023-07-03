FROM golang:alpine

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base && apk add --no-cache curl jq

#swagger
RUN dir=$(mktemp -d) && \
    git clone https://github.com/go-swagger/go-swagger "$dir" && \
    cd "$dir" && \
    go install ./cmd/swagger

WORKDIR /app
COPY go.mod go.sum ./
COPY . .

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -polling -log-prefix=false -build="go build -o apibin ./cmd/api" -command="./apibin" -directory="./"