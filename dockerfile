# Gunakan image Golang sebagai base image
FROM golang:1.23.6-alpine3.21 as builder

# Set working directory
WORKDIR /app

# Copy file Go Modules dan download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy kode aplikasi
COPY . .

# Build aplikasi
RUN go build -o myapp .

# Stage kedua untuk mengurangi ukuran image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy aplikasi yang sudah dibuild dari stage pertama
COPY --from=builder /app/myapp .

# Expose port aplikasi berjalan
EXPOSE 8080

# Jalankan aplikasi
CMD ["./myapp"]
