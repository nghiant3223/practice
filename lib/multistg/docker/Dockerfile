FROM golang:1.14 AS builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

FROM scratch AS production

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]