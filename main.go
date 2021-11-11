// (c) Colin O'Dell <colinodell@gmail.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
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

func init() {
	flag.IntP("port", "p", 8080, "port to listen on")
	flag.Int("code", 200, "response status code to return")
	flag.String("body", "", "response body to return")
	flag.String("headers", "Content-Type: text/plain|Cache-Control: public, max-age=604800", "headers to add to the request (use pipes to separate multiple headers)")
	flag.BoolP("verbose", "v", false, "verbose logging")

	viper.SetEnvPrefix("HTTP")
	viper.AutomaticEnv()

	flag.Parse()
	viper.BindPFlags(flag.CommandLine)

	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Ok"))
	})

	responseCode := viper.GetInt("code")
	responseBody := viper.GetString("body")
	headers := parseHeaders(viper.GetString("headers"))
	port := ":" + strconv.FormatInt(viper.GetInt64("port"), 10)

	handler := rootHandler(responseCode, []byte(responseBody), headers)

	if viper.GetBool("verbose") {
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
