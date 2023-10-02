# syntax=docker/dockerfile:1
# A sample microservice in Go packaged into a container image.
FROM golang:1.19

RUN mkdir /app

ADD . /app

WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download 

RUN go build -o main ./cmd/api

EXPOSE 4000
CMD [ "/app/main" ]
