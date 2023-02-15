# Client IP Server
This is a simple Go program that serves a web endpoint that returns the IP address of the client that made the request. The program accepts a command-line argument to specify the port number to listen on. If the HNIS_PORT environment variable is set, it will override the default port number.

## Getting Started
To run this program on your local machine, you need to have Go installed. You can download and install the latest version of Go from the official website: https://golang.org/doc/install.

1. Clone the repository to your local machine:

```bash
git clone https://github.com/akalp/clientIpServer.git
```

2. Change into the project directory:

```bash
cd clientIpServer
```

3. Run the program using the go run command:

```bash
go run getClientIp.go
```

The program will start serving on the default port 8080. To specify a different port, use the port flag:

```bash
go run getClientIp.go -port=1234
```

This will start the program on port 1234.

If you want to override the default port using an environment variable, you can set the HNIS_PORT variable:

```bash
export HNIS_PORT=5678
go run getClientIp.go
```

This will start the program on port 5678.

### Build

* for linux x64:
```bash
GOOS=linux GOARCH=amd64 go build
```
* for windows x64:
```bash
GOOS=windows GOARCH=amd64 go build
```

### License

This code is released under the MIT License. See LICENSE.md for details.