# static-response-server

[![GitHub](https://img.shields.io/github/license/colinodell/static-response-server?style=flat-square)](https://github.com/colinodell/static-response-server/blob/main/LICENSE)
[![Docker Image Size (latest)](https://img.shields.io/docker/image-size/colinodell/static-response-server?style=flat-square)](https://hub.docker.com/repository/docker/colinodell/static-response-server)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/colinodell/static-response-server?style=flat-square)](https://pkg.go.dev/github.com/colinodell/static-response-server)

Super tiny HTTP server that always returns the same response.

## Purpose

After decommissioning a service, you may want to keep _something_ online and responding to HTTP requests letting users
know the service no longer exists. This is especially useful if legacy code is still hitting an API endpoint and completely
taking that offline might break consumers.

Instead of spinning up a full blown HTTP server like nginx to handle this, you can instead use this super tiny, statically-compiled
Golang-based Docker image which uses as few resources as possible.

## Installation

### Docker

The easiest way to use this server is via Docker:

```bash
docker run -d -p 80:8080 colinodell/static-response-server --code=404 --body="Not Found" --headers="Content-Type: text/plain" -v
```

Or with environment variables:

```bash
docker run -d -p 80:8080 -e HTTP_CODE=404 -e HTTP_BODY="Not Found" -e HTTP_HEADERS="Content-Type: text/plain" -e HTTP_VERBOSE=1 colinodell/static-response-server
```

### Build From Source

Simply clone this project and run `go build` to build the binary.

## Configuration

The server can be configured via command line flags:

```
$ ./static-response-server --help

Usage of ./static-response-server:
--body string      response body to return
--code int         response status code to return (default 200)
--headers string   headers to add to the request (use pipes to separate multiple headers) (default "Content-Type: text/plain|Cache-Control: public, max-age=604800")
-p, --port int         port to listen on (default 8080)
-v, --verbose          verbose logging
```

Environment variables are also supported - simply prefix the flags above with `HTTP_`.

## Examples

### Redirecting all traffic to a different URL

```bash
./static-response-server --body "Moved Permanently" --code 301 --headers "Location: https://www.google.com"
```

### Returning a 404

```bash
./static-response-server --body "This service no longer exists" --code 404
```

### Returning a 404 (using environment variables)

```bash
HTTP_BODY="This service no longer exists" HTTP_CODE=404 ./static-response-server
```

### Pretending your API still accepts POST requests

```bash
./static-response-server --code 201
```
