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
<br /> <br />
## Additional Questions
- Why throttling is useful (if it is)? How would you implement it here? <br />
**_Throttling is useful because it prevents a single client from overwhelming the server with requests, thereby affecting traffic for other clients. It also helps put cost under control. I will implement a simple middleware/interceptor that limits the amount of requests to the server from a particular IP or id (drone) within a time period._**<br /><br />

- What we have to change to make DNS be able to service several sectors at the same
  time? <br />
  **_To make DNS be able to service several sectors, drones would be made to send the sectorId. It would be nice to also spin up more instances of DNS in order to manage traffic adequately._**<br /><br />

- Our CEO wants to establish B2B integration with Mom's Friendly Robot Company by allowing cargo ships of MomCorp to use DNS . The only issue is - MomCorp software
  expects loc value in location field, but math stays the same. How would you approach this? What’s would be your implementation strategy? <br />
  **_I'd simply add another field (location) to my response object, the new field would have the same value as loc._**<br /><br />

- Atlas Corp mathematicians made another breakthrough and now our navigation math is even better and more accurate, so we started producing a new drone model, based on new math. How would you enable scenario where DNS can serve both types of clients? <br />
  **_Based on the new math, I would create a new implementation of the DnsService interface then find a way to route request from the new done to the new interface implementation while I keep the old one._** <br /><br />

- In general, how would you separate technical decision to deploy something from business decision  to release something? <br />
**_This would depend on several factors: time, customers' satisfaction, cost etc. I'd definitely prioritize customers' satisfaction_**
