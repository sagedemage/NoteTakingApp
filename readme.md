# Note Taking App
This is a note taking app. I create notes when learning new technologies and 
remembering what was in my mind lasted time. My goal is is to create a note taking web app 
I can use on my home server. I use raspberry pi as my home server.

## Missing Functionality
It is expected that this app is not ready for production. It is missing important features for it to be complete.
The app does not have authentication for registration, login, and logout for users. 

## Building the Project
Build and Run Using Docker Compose (Option 1)
```
docker-compose build
```
```
docker-compose up
```

Quick Run (Option 2)
```
go run -v cmd/app/*.go
```
***Note***: *Quick run is useful if you want to quickly run the project without docker.*

## Frameworks:
* **gin**
	* [repository](https://github.com/gin-gonic/gin)
	* [website](https://gin-gonic.com/)

## Libraries:
* **gorm**
	* [repository](https://github.com/go-gorm/gorm)
	* [website](https://gorm.io/)

