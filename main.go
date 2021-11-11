// (c) Colin O'Dell <colinodell@gmail.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	flag "github.com/spf13/pflag"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	port         = flag.IntP("port", "p", 8080, "port to listen on")
	responseCode = flag.Int("code", 200, "response status code to return")
	responseBody = flag.String("body", "", "response body to return")
	headers      = flag.StringArray("header", []string{}, "header to add to the request (multiple allowed)")
	verbose      = flag.BoolP("verbose", "v", false, "verbose logging")
)

func main() {
	flag.Parse()
	log.SetFlags(log.Ldate | log.Ltime)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	})

	handler := rootHandler(*responseCode, []byte(*responseBody), *headers)

	if *verbose {
		handler = logRequest(handler)
	}

	http.Handle("/", handler)

	port := ":" + strconv.FormatInt(int64(*port), 10)

	log.Printf("Listening at 0.0.0.0%v", port)
	log.Fatalln(http.ListenAndServe(port, nil))
}

func rootHandler(responseCode int, responseBody []byte, headers []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, header := range headers {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) != 2 {
				http.Error(w, "invalid header", http.StatusBadRequest)
				return
			}

			w.Header().Add(parts[0], strings.Trim(parts[1], " "))
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
