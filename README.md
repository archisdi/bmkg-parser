# BMKG-REST-API

A low latency rest api to retrieve BMKG data to JSON

## Run
To run the app, simply run the following command in the project root directory. 
```bash
bin/run
```

## Build

To build the app, run the following command in the project root directory. 
```bash
bin/setup
```
It will create a binary file required to run the app. After the built process succeeds, it will run a unit test to make sure all functionality runs as expected. 

## Endpoints
All endpoints are defined in ```spec.yaml``` file and can be imported into ```Postman``` collection. Or you can see the detailed [documentation](https://app.swaggerhub.com/apis-docs/archisdi/warpin-messaging/1.0).

## Dockerfile
You can get production ready docker image from [dockerhub](https://hub.docker.com/repository/docker/archisdi/bmkg-rest).
```bash
docker pull archisdi/bmkg-rest
```

## Test
To run the unit test, run the following command in the project root directory. 
```bash
bin/test
```