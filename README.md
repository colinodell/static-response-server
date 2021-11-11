# static-response-server

Super tiny HTTP server that always returns the same response.

## Purpose

After decommissioning a service, you may want to keep _something_ online and responding to HTTP requests letting users
know the service no longer exists. This is especially useful if legacy code is still hitting an API endpoint and completely
taking that offline might break consumers.

Instead of spinning up a full blown HTTP server like nginx to handle this, you can instead use this super tiny, statically-compiled
Golang-based Docker image which uses as few resources as possible.

## Build

Simply clone this project and run `go build` to build the binary.

## Usage

The server can be configured via command line flags:

```
$ ./static-response-server --help

Usage of ./static-response-server:
--body string          response body to return
--code int             response status code to return (default 200)
--header stringArray   header to add to the request (multiple allowed)
-p, --port int             port to listen on (default 8080)
-v, --verbose              verbose logging
```

## Examples

### Redirecting all traffic to a different URL

```bash
./static-response-server --body "Moved Permanently" --code 301 --header "Location: https://www.google.com"
```

### Returning a 404

```bash
./static-response-server --body "This service no longer exists" --code 404
```

### Pretending your API still accepts POST requests

```bash
./static-response-server --code 201
```