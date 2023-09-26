# Use the official Golang image as a build environment.
FROM golang:1.17 AS build

# Set the working directory inside the container.
WORKDIR /app

# Copy go module files and download dependencies.
COPY ./* ./

RUN go mod download

# Copy the entire source code from the current directory to the working directory inside the container.
#COPY . .

# Build the Go application with the desired output name.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wait-for-db .

