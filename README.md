# go-demo

This project aims to be a first try to get a simple go REST server running. 
Use gin as dependency for REST: https://github.com/gin-gonic/gin#quick-start

## Learnings so far
- Set GOPATH correctly: https://github.com/golang/go/wiki/SettingGOPATH It is also import to understand how go and go projects are structure
- Tried to use dep as dependency manager. With the rise of go modules, go modules seem the way to go. Used it to install gin for setting up REST interfaces:
``
go get -v github.com/gin-gonic/gin
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
