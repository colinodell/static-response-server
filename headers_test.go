// (c) Colin O'Dell <colinodell@gmail.com>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseEmptyHeaders(t *testing.T) {
	headers := parseHeaders("")
	assert.Equal(t, 0, len(headers))
}

func TestParseSingleHeader(t *testing.T) {
	headers := parseHeaders("Content-Type: text/plain")
	assert.Equal(t, 1, len(headers))
	assert.Equal(t, "Content-Type", headers[0].name)
	assert.Equal(t, "text/plain", headers[0].value)
}

func TestParseMultipleHeaders(t *testing.T) {
	headers := parseHeaders("Content-Type: text/plain | Content-Length: 123")
	assert.Equal(t, 2, len(headers))
	assert.Equal(t, "Content-Type", headers[0].name)
	assert.Equal(t, "text/plain", headers[0].value)
	assert.Equal(t, "Content-Length", headers[1].name)
	assert.Equal(t, "123", headers[1].value)
}

func TestParseHeaderWithColon(t *testing.T) {
    headers := parseHeaders("Location: https://www.example.com")
    assert.Equal(t, 1, len(headers))
	assert.Equal(t, "Location", headers[0].name)
	assert.Equal(t, "https://www.example.com", headers[0].value)
}