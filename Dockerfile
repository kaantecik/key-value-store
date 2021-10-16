FROM golang:1.16-alpine3.14 as builder

COPY . .

RUN go run .