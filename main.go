// (c) Colin O'Dell <colinodell@gmail.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type header struct {
	name, value string
}

func parseHeaders(headers string) []header {
	var result []header

    // split string into array of headers
	arr := strings.Split(headers, "|")
	for _, v := range arr {
        // split header into name and value
        hdr := strings.Split(v, ":")
        if len(hdr) != 2 {
            log.Fatal("Invalid header: " + v)
        }

        result = append(result, header{strings.Trim(hdr[0], " "),  strings.Trim(hdr[1], " ")})
    }

	return result
}

var (
	port         = kingpin.Flag("port", "Port to listen on").Short('p').Default("8080").Envar("HTTP_PORT").Int64()
	headers      = kingpin.Flag("headers", "Headers to add to the response").Default("").Envar("HTTP_HEADERS").String()
	responseCode = kingpin.Flag("code", "HTTP status code to return").Default("200").Envar("HTTP_CODE").Int()
	responseBody = kingpin.Flag("body", "Body to return").Default("").Envar("HTTP_BODY").String()
	verbose      = kingpin.Flag("verbose", "Verbose logging").Short('v').Envar("VERBOSE").Bool()
)

func main() {
	kingpin.Parse()

	log.SetFlags(log.Ldate | log.Ltime)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	})

	headers := parseHeaders(*headers)
	port := ":" + strconv.FormatInt(*port, 10)

	handler := rootHandler(*responseCode, []byte(*responseBody), headers)

	if *verbose {
		handler = logRequest(handler)
	}

	http.Handle("/", handler)

	log.Printf("Listening at 0.0.0.0%v", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}

func rootHandler(responseCode int, responseBody []byte, headers []header) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, header := range headers {
			w.Header().Add(header.name, header.value)
		}

		w.WriteHeader(responseCode)
		w.Write(responseBody)
	})
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s -- %s %s\n", r.Method, r.URL, r.RemoteAddr, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}
