# Golang gRPC with REST API

## Introduction
here's making gRPC with REST API gateway with echo Framework (Go Programming Language).

## Installation
#### 1. Requirements

##### a. Golang Language SDK minimal 1.13 https://golang.org/dl/
##### b. Dep Golang Package Management https://golang.github.io/dep/docs/installation.html

#### 2.Clone The Projects
Go to your src/ folder inside go project directory
```$xslt
    cd ~YOUR/GO/FOLDER/DIRECTORY/src/
``` 
then clone the project
```$xslt
    git clone https://github.com/huzairuje/grpc_golang.git
```
#### 3. install all the dependencies
a. Go to the project folder
```$xslt
    cd ~YOUR/GO/FOLDER/DIRECTORY/src/grpc_golang/
```
b. then run `dep ensure` command.
```
```
#### 4. Run The Server
a. Go to the `service-garage` folder
```$xslt
    cd ~YOUR/GO/FOLDER/DIRECTORY/src/grpc_golang/services/service-garage
```
b. then run `go run main.go` command.
```$xslt
    go run main.go
```
c. Go to the `service-user` folder
   ```$xslt
       cd ~YOUR/GO/FOLDER/DIRECTORY/src/grpc_golang/services/service-user
   ```
d. then run `go run main.go` command.
   ```$xslt
       go run main.go
   ```
#### 5. Run The Client
a. Go to The Client folder
```$xslt
    cd ~YOUR/GO/FOLDER/DIRECTORY/src/grpc_golang/client
```
b. and run the client
 ```$xslt
     go run main.go
 ```