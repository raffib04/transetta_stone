# Build stage
FROM golang:1.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# Run stage
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/app /app/app
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/data /app/data 
COPY .env .env  
EXPOSE 8080
CMD ["./app"]
