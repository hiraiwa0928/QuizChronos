FROM golang:1.22.3

WORKDIR /app

# Copy files
COPY ./src .

# Download all dependencies
RUN go mod download

# Build the application
RUN GOARCH=amd64 GOOS=linux go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

