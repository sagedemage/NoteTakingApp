# syntax=docker/dockerfile:1

# Use Debian as OS
FROM golang:1.19-bullseye

# Update and install gcc and make
RUN apt update && \
	apt install gcc

# Set Working Directory of inside the container
WORKDIR /app

# Copy files from local to the working directory
ADD . ./

# Copy the go source files from the cmd/app directory to the working directory
ADD cmd/app/*.go ./

# Install the go external packages
RUN go get -u github.com/gin-gonic/gin \ 
	gorm.io/gorm \ 
	gorm.io/driver/mysql \ 
	github.com/gin-contrib/cors \ 
	github.com/golang-jwt/jwt/v4 \
	github.com/joho/godotenv

# Build the Go app
RUN go build -o build/ -v ./cmd/app/...

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD [ "./build/app" ]

