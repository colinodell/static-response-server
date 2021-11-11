// (c) Colin O'Dell <colinodell@gmail.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"log"
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
		// skip any empty lines
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		// split header into name and value
		hdr := strings.Split(v, ":")
		if len(hdr) != 2 {
			log.Fatal("Invalid header: " + v)
		}

		result = append(result, header{strings.TrimSpace(hdr[0]), strings.TrimSpace(hdr[1])})
	}

	return result
}
