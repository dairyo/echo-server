# syntax=docker/dockerfile:1
FROM golang:1.18.4-alpine3.16 AS build
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build

FROM alpine:latest
COPY --from=build /app/echo-server ./
CMD ["./echo-server"]

