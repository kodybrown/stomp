package stomp

// http://stomp.github.io/stomp-specification-1.2.html
// Stomp message content:
//
//    COMMAND
//    org-id:111
//    content-type:text/json
//
//    {"key":"value"}^@
//

const (
	EOM     = "^@" // end of message
	NewLine = "\n"
)
