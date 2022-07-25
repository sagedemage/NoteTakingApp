# syntax=docker/dockerfile:1

# Use Alpine as OS
FROM golang:1.16-bullseye

WORKDIR /app

ADD . /app

RUN apt update
RUN apt install gcc make

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go env -w GO111MODULE=on

RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite

COPY cmd/app/*.go ./

RUN go build -o out

EXPOSE 8080

CMD [ "./out" ]





