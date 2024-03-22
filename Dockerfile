FROM golang:1.22.1-alpine3.19

WORKDIR /usr/src/app
RUN apk add inkscape
COPY go.mod go.sum ./
COPY . .
RUN go mod download && go mod verify
RUN go build -v -o /usr/local/bin ./...

EXPOSE 80

ENTRYPOINT "convertSvg"