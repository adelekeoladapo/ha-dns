# Drone Navigation Service (DNS)
This project is a simple API that helps drones locate databank to upload data they gather during space exploration.

It's a Golang project, it was developed using Echo (https://echo.labstack.com), a high performance, extensible, minimalist Go web framework.

## Requirements
- Docker
- Postman

## How to run
### Method 1
- Clone or download the project
- Navigate to the root directory on your terminal
- Run `docker build -t ha-dns .` to build the docker image
- Run `docker run -it -p 8000:7001 ha-dns` to run the generated image in a container. 
### Method 2
- Clone or download the project
- Navigate to the root directory on your terminal
- Run `make docker-build` to build the docker image
- Run `make docker-run` to run the generated image in a container.

## How to use DNS
The DNS service runs in an iterative mode on port `8000` of the host machine/computer. Kindly use Postman to test the service. <br />
It exposes a GET endpoint with the details below: <br />
URL: `http://localhost:8000/api/v1/dns` <br />
JSON Request Sample: `{
  "x": "123.12","y": "456.56",
  "z": "789.89",
  "vel": "20.0",
  }`