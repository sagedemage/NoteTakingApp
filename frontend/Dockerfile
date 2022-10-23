# syntax=docker/dockerfile:1

# Use Alpine as OS
FROM node:16-alpine

# Update and install gcc and make
RUN apk update && \
	apk add gcc && \
	apk add --update nodejs npm

# Set Working Directory of inside the container
WORKDIR /app

# Copy app files 
COPY . .

# Install all dependencies
RUN npm install react react-dom react-router-dom react-bootstrap

# Install serve
RUN npm install -g serve

# Build app
RUN npm run build

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD [ "serve", "-s", "build" ]

