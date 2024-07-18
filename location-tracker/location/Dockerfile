FROM golang:1.22-alpine3.19 as builder

WORKDIR /app

RUN apk --no-cache update && \
apk --no-cache add git gcc libc-dev

COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o /processor cmd/processor/main.go
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags musl -o /server cmd/server/main.go


FROM alpine:latest AS server

COPY --from=builder /server /server
COPY --from=builder /processor /processor


