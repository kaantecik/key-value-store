### Stage 1
FROM golang:1.16-alpine AS builder

WORKDIR /tmp/key-value-store

# We want to populate the module cache based on the go.{mod,sum} files.
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /bin/main /tmp/key-value-store/cmd/key-value-store

### Stage 2
FROM alpine AS production

RUN apk --update --no-cache add bash ca-certificates

COPY --from=builder /bin/main /bin/main

ENTRYPOINT ["/bin/main"]
