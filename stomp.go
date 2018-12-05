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

// /// <summary>
// /// Gets a Stomp message for not authorized.
// /// </summary>
// public readonly static Message NotAuthorized = new Message() {
// 	Command = Commands.Nack,
// 	Headers = new Bricksoft.Stomp.Headers(),
// 	Body = "NOT AUTHORIZED^@",
// };

// /// <summary>
// /// Gets a Stomp message for an unknown command.
// /// </summary>
// public readonly static Message UnknownCommand = new Message() {
// 	Command = Commands.Error,
// 	Headers = new Bricksoft.Stomp.Headers(),
// 	Body = "UNKNOWN COMMAND^@",
// };
