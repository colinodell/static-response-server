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

Using Docker-Compose? We've got that covered too!

```yaml
version: '3'
services:
  static-response-server:
    image: colinodell/static-response-server
    ports:
      - "80:8080"
    environment:
      - HTTP_CODE=404
      - HTTP_BODY="Not Found"
      - HTTP_HEADERS="Content-Type: text/plain"
      - HTTP_VERBOSE=1
```

_(Consider using a reverse proxy like [Traefik](https://github.com/traefik/traefik) to secure the requests with HTTPS.)_

### Build From Source

Simply clone this project and run `go build` to build the binary.

## Configuration

The server can be configured via command line flags or environment variables:

| Flag                | Environment Variable | Default | Description                                                        |
|---------------------|----------------------|---------|--------------------------------------------------------------------|
| `--port` or `-p`    | `HTTP_PORT`          | `8080`  | Port to listen on                                                  |
| `--code`            | `HTTP_CODE`          | `200`   | HTTP status code to return                                         |
| `--body`            | `HTTP_BODY`          |         | HTTP body to return                                                |
| `--headers`         | `HTTP_HEADERS`       |         | HTTP headers to return (multiple headers separated by pipes (`\|`) |
| `--verbose` or `-v` | `VERBOSE`            | (off)   | Print verbose output                                               |

```
$ ./static-response-server --help

usage: static-response-server [<flags>]

    Flags:
    --help           Show context-sensitive help (also try --help-long and --help-man).
    -p, --port=8080  Port to listen on
    --headers=""     Headers to add to the response
    --code=200       HTTP status code to return
    --body=""        Body to return
    -v, --verbose    Verbose logging
```

## Examples

### Returning a 404

```bash
./static-response-server --body "This service no longer exists" --code 404
```

### Returning a 404 (using environment variables)

```bash
HTTP_BODY="This service no longer exists" HTTP_CODE=404 ./static-response-server
```

### Redirecting all traffic to a different URL

```bash
./static-response-server --body "Moved Permanently" --code 301 --headers "Location: https://www.google.com"
```

### Pretending your API still accepts POST requests

```bash
./static-response-server --code 201
```
