# Backend

## Build and Run the App with Docker
Build the app
```
docker-compose build
```

Startup the app
```
docker-compose up
```

## Build and Run the App without Docker

## Dependencies
Install dependencies
```
go get -u github.com/gin-gonic/gin \ 
	gorm.io/gorm \
	gorm.io/driver/sqlite \
	github.com/golang-jwt/jwt/v4 \
	github.com/joho/godotenv
```

## Mage build system Installation
Install Mage build system
```
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```
If on Linux, export the go bin direcotory
```
export PATH="$PATH:$HOME/go/bin/"
```

## Run web app
```
mage -v run
```

## Test the web app
```
mage -v test
```

