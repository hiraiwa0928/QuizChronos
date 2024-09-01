FROM golang:1.22.3

WORKDIR /app

# Copy files
COPY ./src .

# Download all dependencies
RUN go mod download

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["go", "build", "-o", "main", ".", "&&", "./main"]
