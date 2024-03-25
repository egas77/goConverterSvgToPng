FROM golang:1.22.1-alpine3.19 AS prepare
WORKDIR /usr/src/app
COPY . .
RUN go mod download && go mod verify
RUN go build convertSvg

FROM alpine:3.19 AS api
RUN apk add inkscape
WORKDIR /usr/local/bin
COPY --from=prepare /usr/src/app/convertSvg .
EXPOSE 80
ENTRYPOINT "convertSvg"