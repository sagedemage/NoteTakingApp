# Backend

## Dependencies
Install dependencies
```
go get -u github.com/gin-gonic/gin \ 
	gorm.io/gorm \
	gorm.io/driver/sqlite
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
