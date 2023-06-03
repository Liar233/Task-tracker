FROM golang:1.19.4 AS build-server

WORKDIR /app

COPY . .

RUN go mod download \
    && go mod verify \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /output/tracker /app/cmd/app/

FROM alpine:latest

WORKDIR /app

COPY --from=build-server /output/tracker /app/tracker

RUN apk add tzdata

RUN chmod 755 /app/tracker
RUN chmod +x /app/tracker

EXPOSE 8080

ENTRYPOINT ["/app/tracker"]