# Build stage.
FROM golang:1.14.4 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/server .

# Final stage.
FROM alpine:3.12.0
WORKDIR /app

ENV DATABASE_DIALECT=postgres \
    DATABASE_URL= \
    REDIS_URL= \
    JWT_SECRET=defaultsecretkey

COPY --from=builder /bin/server .
CMD ["./server"]
