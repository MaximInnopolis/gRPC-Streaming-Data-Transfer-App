# gRPC Streaming Data Transfer App

This is a console-based gRPC streaming data transfer application that includes both a server and a client.

## Features

- The server continuously streams incrementing numbers to the client.
- The client sends a request to the server to start streaming, specifying the interval in milliseconds.
- The client saves received numbers in a buffer and prints them to the console after a certain duration or when a specific buffer size is reached.
- The client can send a request to the server to stop the data transmission.

## Usage

### Server

To start the server, run the following command:

```bash
go run server.go
```

### Server Flags
You can customize the client behavior by using the following optional flags:

1. -help: Prints all  flag information (e.g., -help)

2. -Port: Specifies the Port to listen on (e.g., -Port "port)

Example usage with custom flags:
```bash
go run client.go -Login "john_doe" -Password "secret" -IntervalMs 1000 -BufferSize 20 -StopAfter 10s
```


### Client

To start client , run the following command:

```bash
go run client.go -Login "user" -Password "pass"
```

Note that the parameters **IntervalMs**, **BufferSize**, and **StopAfter**
have initial values of 500, 10, and 5 seconds, respectively.

### Client Flags
You can customize the client behavior by using the following optional flags:

1. -help: Prints all  flag information (e.g., -help)

2. -Login: Specifies the client login. (e.g., -Login "user")

3. -Password: Specifies the client password. (e.g., -Password "pass")

4. -IntervalMs: Specifies the interval in milliseconds between received data points. (default: 500)

5. -BufferSize: Specifies the maximum buffer size for received data points. (default: 10)

6. -StopAfter: Specifies the duration after which the client will stop data reception. (default: 5 seconds)

Example usage with custom flags:
```bash
go run client.go -Login "john_doe" -Password "secret" -IntervalMs 1000 -BufferSize 20 -StopAfter 10s
```
