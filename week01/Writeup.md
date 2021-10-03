I graduated with a bachelor's degree in Information Technology from Andhra University, India. Started career in implementing network router and 4G Enodeb modules using C and C++ languages and then joined a startup to work on SDN and NFV applications which were developed in Java and Python languages. Submitted OneConvergence network plugin to
Openstack Neutron component. Contributed to OpenContrail SDN components which were in C++ language. Currently developing network analytics applications based on microservices architecture using 
Spring Boot Java framework and deploying in Kubernetes and Azure public cloud. Efficient in writing algorithms using the right data structures by following SOLID and design principles. Key strengths are computer science foundations and debugging the code

 
I am fed up with developing applications using frameworks rather than libraries. So decided to develop web applications and cloud-native microservices using GO language. The reason behind picking GO language is due to standard libraries, inbuilt tools, concurrency patterns, performance, and focusing more time to learn GO foundations and writing idiomatic GO code using the standard library and writing test cases using Test Driven Development [TDD].

I want to develop customize command-line tools for existing and future applications. 

My tech stack is mentioned below 

- Language : GO, GO-Kit

- Communication : gRPC, Protocol Buffers        

- Messaging systems : NATS, NATS Streaming, NATS Jetstream

- Database : Cockroach DB

- Tracing : Zipkin

- ServiceMesh : Istio

- Observeability : OpenTelemetry

- Orchestrator : Kubernetes

- Architecture : Microservices, Event-Driven Streaming Architecture, DDD, Event Sourcing, CQRS, Clean Architecture

All the above software is developed using GO Language only. I will be completely moving to GO Language as a full-time GO software developer.
Going forward will be working based on cloud-native application architecture and cloud-native infrastructures using a cloud-native tech stack.

fmt stands for the Format package. This package allows to format basic strings, values, or anything and print them or collect user input from the console, or write into a file using a writer, or even print customized fancy error messages. This package is all about formatting input and output. Exploring fmt package based on the URL (https://pkg.go.dev/fmt) gave a lot of information regarding functions with "example". Clear distinctions and sample examples among verbs (%v, %#v, %q, %s) help developers to fulfill their requirements. %T will give the type of the variable. fmt.Print, fmt.Printf, fmt.Println are some of the functions.
 

The string interface and GoString interface helps in overriding the default behavior. Stringer is implemented by any value that has a String method, which defines the “native” format for that value.
The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted printer such as Print. Implementing Stringer is useful for many purposes, such as for logging and debugging. GoStringer is implemented by any value that has a GoString method, which defines the Go syntax for that value. The GoString method is used to print values passed as an operand to a %#v format. Implemented sample example by overriding the GoString() function and implemented sample example by overriding the String() function. 

Below are the differences between Go Lang and Java print functions :

```sh
GO :

    func Print(a ...interface{}) (n int, err error)

    func Printf(format string, a ...interface{}) (n int, err error)

    func Println(a ...interface{}) (n int, err error)
```
```sh
JAVA:

    public void print(char c)
    
    public void println(int x)
```

 

There is no error return type for print functions in Java and need to implement for every type.