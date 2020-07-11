# go-demo

This project aims to be a first try to get a simple go REST server running. 
Use gin as dependency for REST: https://github.com/gin-gonic/gin#quick-start

## Learnings so far
- Set GOPATH correctly: https://github.com/golang/go/wiki/SettingGOPATH It is also import to understand how go and go projects are structure
- Tried to use dep as dependency manager. With the rise of go modules, go modules seem the way to go. Used it to install gin for setting up REST interfaces:
``
go mod init // Create a module so we can use the go module system
``

## Test REST
Perform GET to ping interface
```
curl localhost:8080/ping
```
Expected result:
```
{"message":"pong"}
```

## Build docker image
Provided is a dockerfile to build a distroless image of the application:
```
docker build -t go-demo:0.0.1 .
```

Run image:
```
docker run -d -p 8081:8080 go-demo:0.0.1 // Run it on port 8081 in case you still have it running locally
```

Test image:
```
curl localhost:8080/ping
```
