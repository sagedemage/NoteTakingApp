# syntax=docker/dockerfile:1

# Use Alpine as OS
FROM node:19-bullseye

# Update and install gcc and make
RUN apt update && \
	apt install -y build-essential nodejs npm

# Set Working Directory of inside the container
WORKDIR /app

# Copy app files 
COPY . .

# Install all dependencies
RUN npm install react react-dom react-router-dom react-bootstrap axios universal-cookie

# Install serve
RUN npm install -g serve

# Build app
RUN npm run build

# Expose port 8080 to the outside world
EXPOSE 3000

# Run the executable
CMD [ "serve", "-s", "build" ]

