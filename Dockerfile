FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o main .

FROM gcr.io/distroless/base-debian12

COPY --from=builder /app/main /app/main

CMD ["./main"]

