# This is a project for learning how to work with gRPC

## 1. Create a gRPC server

### 1.1.Create a .proto file with the Messages and Service definition

* Service: is the inferface that **exposes the methods** that can be called remotelly

* Messages: are the **structure of the payload** that are used as parameters and responses by a method 

### 1.2.Execute the protoc compiler

* Run the compiler: this action will generate the language specific code base on the .proto file

* File .pb: this file has the language specific structures created using the **.proto Messages definitions**

* File _grpc.pb: this file has the language specific interface created using the **.proto Service definition**

### 1.3.Implement the server interface

* In the _grpc.pb file there is an inteface with name **ProtoServiceNameServer**, where **ProtoServiceName** is the name of the service defined in the .proto file. 

* This interface must be implemented and is used by grpc Server for exposing the methods for clients

### 1.4.Create a server

* Create a instance of a grpc server that uses the ProtoServiceNameServer inferface to expose the service

* In go you must:
    * create a **service** instance of the ProtoServiceNameServer interface (that you implemented)
    * create a **server** with grpc.NewServer
    * register the service using the server and the service with the function **RegisterProtoServiceNameServer** (in the file _grpc.pb), which takes both **grpc server** and the **service interface** as arguments
    * run the server (server.Serve)

## 2.Create a gRPC client

### 2.1.Create a client

* use the function **NewProtoServiceNameClient** (in the file _grpc.pb)

* use this client for calling the methods of the Service interface


## 3.Observation

* After generating the .pb files using the protoc compiler (_grpc.pb and .pb) both, client and server, must have access to them (for using the interfaces)