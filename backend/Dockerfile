# Build stage.
FROM golang:1.14.4 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/server .


# Final stage.
FROM alpine:3.12.0
WORKDIR /app

ENV SERVER_PORT=:8080 \
    DB_DIALECT=postgres \
    DB_HOST=db \
    DB_PORT=5432 \
    DB_USERNAME=postgres \
    DB_PASSWORD=password \
    DB_NAME=yarc \
    RDB_ADDR=redis:6379 \
    RDB_PASSWORD= \
    RDB_DB_NUMBER=0 \
    JWT_SECRET=defaultsecretkey \
    FRONTEND_ADDR=http://localhost:8081

COPY --from=builder /bin/server .
EXPOSE 8080
CMD ["./server"]