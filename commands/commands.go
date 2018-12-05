package commands

const (
	Abort       string = "ABORT"
	Ack         string = "ACK"
	Begin       string = "BEGIN"
	Commit      string = "COMMIT"
	Connect     string = "CONNECT"
	Connected   string = "CONNECTED"
	Disconnect  string = "DISCONNECT"
	Error       string = "ERROR"
	Message     string = "MESSAGE"
	Nack        string = "NACK"
	Receipt     string = "RECEIPT"
	Send        string = "SEND"
	Subscribe   string = "SUBSCRIBE"
	Success     string = "SUCCESS"
	Unsubscribe string = "UNSUBSCRIBE"

	// custom commands
	Ping             string = "PING"
	Pong             string = "PONG"
	GetLastMeterRead string = "GET-LAST-METER-READ"
	Read             string = "READ"
)
