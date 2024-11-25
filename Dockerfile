FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o binary

ENTRYPOINT ["/app/binary"]
