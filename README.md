# Ads Service

## Description
This service exposes an HTTP API that serves endpoints to interact with user advertisements. For now this service can 
be used to retrieve the ads that are available for a user depending on the user's location, language, time and country.

## Requirements
* Docker (Optional)
* Golang 1.19^

## Getting Started

As the first step you should clone the repository and then deploy the service. You have two options: with or without docker.

### Run using docker

* To deploy this microservice:
```bash
docker-compose up --build
```

Check that the config file located in `ads-service/config/config.yaml` the redis section looks like this:

```
redis:
  address: "host.docker.internal:6379"
```

### Run without docker

Check that the config file located in `ads-service/config/config.yaml` the redis section looks like this:

```
redis:
  address: "localhost:6379"
```

Or is pointing to the redis instance that you want to point to. 

Then from teh project root you can run the command 

```
go run main.go
```

## Using service

### Get advertisement endpoint
To use the get ad endpoint use the following example as guide:

```
curl --location --request GET 'localhost:8080/v1/advertisement?userId=115127&country=us&language=eng' \
--header 'X-API-Key: My-Ap1-k3y'
```
