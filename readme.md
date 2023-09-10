# Coda Payments Take-home Assignment

## Prerequisites

Before you begin, ensure you have Go installed on your system:

https://go.dev/doc/install

## Simple API

This is an API which accepts an HTTP POST with JSON payload and respond with a success response containing an exact copy of the JSON request it received.

#### Sample Payload

```
{
    "game":"Valorant",
    "gamerID":"Heckermann",
    "points":2500
}
```

### Installation

Clone the repository and go the SimpleAPI folder:

```
cd SimpleAPI
```

Run the API (If you don't add 'PORT' to the command, the default is 8080.)

```
PORT=8081 go run main.go
```

or run the API using makefile

```
make run-api-1
make run-api-2
make run-api-3
```

You can also simulate slow HTTP response using 'slow' tag is second/

```
PORT=8080 slow=10 go run main.go
```

or use makefile (default is 6 second)

```
make run-slow-api-1
```

## Round Robin API

ROUND ROBIN API which receive the response from the application API and send it back to the client. The payload is the same as Simple API. The default Server for Round Robin to forwards HTTP requests from a client to a selected backend server are

```
backendServers := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}
```

go the RoundRobinAPI folder:

```
cd RoundRobinAPI
```

Run the API (If you don't add 'PORT' to the command, the default is 9090.)

```
PORT=9091 go run main.go
```

or run the API using makefile

```
make run-rr
```

### Round Robin Implementation

Round Robin is aload balancing algorithm that distributes incoming requests among a group of backend servers or resources. In this case, the algorithm will distributes the request from port 8080 to 8081 to 8082 and back to 8080 again.

If one of the server is down. The API also have a health check to validate the state of the server and won't distribute the request to that server. If the server is start to work again, Round Robin algorithm will back to distribute the request again.

If one of the server is slow (Slow response time), the API will mark the state as "IsSlow = true" and won't include in the Round Robin process.

Round Robin run on the normal (fast) server first but if there is no fast server running but has only the slow one, Then the API will apply Round Robin to the slow API.It best to return slow response than no response and if the response is now normal. The API will reset the state.

If there is no healthy backend available, it will return error message.
