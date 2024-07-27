FROM golang:alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./employee-service cmd/main.go

FROM alpine:latest AS employees-service
WORKDIR /app
COPY --from=builder /app/employee-service .
EXPOSE 8080
ENTRYPOINT ["./employee-service"]
