FROM golang:1.19 AS builder

COPY . /goproject
WORKDIR /goproject

ENV CGO_ENABLED=1
ENV GOOS=linux

RUN go mod tidy \
 && go build -a -ldflags '-linkmode external -extldflags "-static"' \
    -o /goproject/bin/server /goproject/cmd/server


FROM alpine:latest
COPY --from=builder /goproject/bin /app

RUN sh -c "apk add --update sqlite bash"

ENV CONFIG_FILE /app/config.yaml
WORKDIR /app
COPY ./migrations /migrations
VOLUME /data

ENTRYPOINT ["/app/server"]
