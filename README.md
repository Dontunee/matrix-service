## Matrix Service

To run the application:


1. Run ```go mod tidy``` to ensure the tesify dependency is pulled


2. Run with the command;

```bash
go run cmd/server/main.go
```

To make request, send an HTTP request to the specific endpoint.

For example using curl for the echo route:

```bash
curl -F 'file=@data/matrix.csv' "localhost:8080/echo"
```

### **Testing**

To run test use ``` go test ./.../ ```

You can check test coverage with ```go test -cover ./.../```.
