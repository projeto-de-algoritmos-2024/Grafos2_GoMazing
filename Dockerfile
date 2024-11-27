# Use the official Golang image as the base image
FROM golang:1.21

# Install necessary libraries for OpenGL, X11, and xauth
RUN apt-get update && apt-get install -y \
    libgl1-mesa-dev \
    xorg-dev \
    xauth \
    && rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o gomazing main.go

# Set the environment variable for the X11 display
ENV DISPLAY=${DISPLAY}

# Expose the port the application uses
EXPOSE 8080

# Command to run the application
CMD ["./gomazing"]
