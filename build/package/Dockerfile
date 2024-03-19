# Builder stage
FROM --platform=arm64 golang:1.22-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/lite/lite.go

# Final stage
FROM --platform=arm64 alpine:3.14

# Consider adding a non-root user for running the application
# RUN adduser -D myuser
# USER myuser

WORKDIR /app

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
