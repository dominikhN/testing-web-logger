# testing-web-logger

This is a simple tool to spawn a webserver that will log all requests (date, remote IP, HTTP method, and path) to stdout. Especially usefol for testing container setups.

Upon a request against */health*, it will respond with status 200 and a body of *OK*.
If a request against any other path is performed, it will send a redirect (302) to */health*.

## Building

```bash
go get -d ./...
go build
```

## Running

This program will listen on port 8443 by default:

```bash
./testing-web-logger
```

You can specify an alternative port by supplying an `-port` argument:

```bash
./testing-web-logger -port 8080
```
