FROM golang:1.22 AS builder

WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the rest of the application code
# This assumes your Go application is in the current directory

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/main /app/main

CMD ["./main"]

