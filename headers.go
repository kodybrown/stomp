package stomp

import (
	"fmt"
	"strings"
)

// Headers ..
type Headers map[string]string

// Get ..
func (h Headers) Get(key string) string {
	var value, exists = h[key]
	if !exists {
		return ""
	}
	return value
}

// String ..
func (h *Headers) String() string {
	var s = ""

	for key, value := range *h {
		s += fmt.Sprintf("%s%s:%s", NewLine, key, value)
	}

	return strings.TrimLeft(s, NewLine)
}
