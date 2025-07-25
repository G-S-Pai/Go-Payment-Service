# Build Stage
FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-payment-service .

# Minimal runtime image
EXPOSE 8080
CMD ["./go-payment-service"]