# Notebook
[![Production Workflow](https://github.com/sagedemage/NotebookApp/actions/workflows/prod.yml/badge.svg)](https://github.com/sagedemage/NotebookApp/actions/workflows/prod.yml)

[![Development Workflow](https://github.com/sagedemage/NotebookApp/actions/workflows/dev.yml/badge.svg)](https://github.com/sagedemage/NotebookApp/actions/workflows/dev.yml)

![](images/dashboard-page.webp)

## Purpose
This is a note taking app. I create notes when learning new technologies and 
remembering what was in my mind the last time. My goal is is to create a note taking web app 
I can host on a web server.

## Building the Project
Go to the root of the repository
```
cd Notebook
```

Build the Docker image
```
docker-compose build
```

Start up the Docker image
```
docker-compose up
```

## Frontend:
### Frameworks:
* [facebook/react](https://github.com/facebook/react/)
	* [reactjs.org](https://reactjs.org/)

### Libraries:
* [axios/axios](https://github.com/axios/axios)
* [universal-cookie](https://github.com/reactivestack/cookies/tree/master/packages/universal-cookie)

## Backend:
### Frameworks:
* [gin-gonic/gin](https://github.com/gin-gonic/gin)
	* [gin-gonic.com](https://gin-gonic.com/)

### Libraries:
* [go-gorm/gorm](https://github.com/go-gorm/gorm)
	* [gorm.io](https://gorm.io/)
* [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
* [joho/godotenv](https://github.com/joho/godotenv)


