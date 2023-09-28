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

### Client

To start client , run the following command:

```bash
go run client.go -Login "user" -Password "pass"
```

Note that the parameters **IntervalMs**, **BufferSize**, and **StopAfter**
have initial values of 500, 10, and 5 seconds, respectively.