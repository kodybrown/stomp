package stomp

import (
	"fmt"
	"strings"
)

// A Message represents a message received from the STOMP server.
// In most cases a message corresponds to a single STOMP MESSAGE frame
// received from the STOMP server.
type Message struct {
	// Command of the message.
	command string

	// Optional header entries.
	// These are the header entries received with the message.
	headers Headers

	// The message body, which is an arbitrary sequence of bytes.
	// The ContentType indicates the format of this body.
	body []byte

	// MIME content type.
	contentType string // MIME content
}

// Command ..
func (m *Message) Command() string {
	return m.command
}

// SetCommand ..
func (m *Message) SetCommand(command string) {
	m.command = command
}

// Body ..
func (m *Message) Body() string {
	return string(m.body)
}

// SetBody ..
func (m *Message) SetBody(body string) {
	m.body = []byte(body)
}

// Headers ..
func (m *Message) Headers() Headers {
	return m.headers
}

// AddHeader ..
func (m *Message) AddHeader(key, value string) {
	m.headers[key] = value
}

// String ..
func (m *Message) String() string {
	return fmt.Sprintf("%s\n%s\n\n%s%s", m.command, m.headers.String(), string(m.body), EOM)
}

// NewMessage creates a new Message.
func NewMessage() *Message {
	var m = Message{}

	m.command = ""

	m.body = []byte("")
	m.contentType = "text/plain"

	m.headers = Headers{}

	return &m
}

// ParseMessage ..
func ParseMessage(msg string) (*Message, error) {
	var message = NewMessage()

	msg = strings.TrimSpace(msg)

	if !strings.HasSuffix(msg, EOM) {
		return nil, fmt.Errorf("Invalid message format: missing `%s` at end of body", EOM)
	}

	msg = strings.TrimRight(msg, EOM)

	var lines = strings.Split(msg, NewLine)
	var isBody = false
	var body = ""

	if len(lines) < 2 {
		return nil, fmt.Errorf("Invalid message format: not enough lines")
	}

	var cmd string

	// for (int i = 1; i < lines.Length; i++) {
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if cmd == "" {
			// The line is the command..
			// log.Println("command =", line)
			cmd = line
			continue
		}

		if isBody {
			// log.Println("body +=", line)
			body += NewLine + line
		} else {
			if line == "" {
				// the rest is the body.
				// log.Println("switch to body")
				isBody = true
				continue
			}

			// The line is a header..
			// log.Println("header  =", line)
			var header = strings.SplitN(line, ":", 2)
			message.Headers()[header[0]] = header[1]
		}
	}

	message.SetCommand(cmd)

	// Remove first NewLine (that was added above)
	if len(body) > len(NewLine) {
		message.SetBody(body[len(NewLine):])
	}

	return message, nil
}
