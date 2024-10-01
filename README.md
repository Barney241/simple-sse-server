# Simple SSE server for testing

This repo contains simple implementation of SSE server and client in go.

## Usage

start server
```bash
cd server && go run main.go
```
after server is running start client
```bash
cd client && go run main.go
```
Default scenerio looks like this. Server will sent 10 messages to connected client and then close the connection. This behavior or data format can be easily changed in `server/main.go`

## Testing with different clients
If you are not using go u can just call `GET http://localhost:8888/events` in any language or tool
