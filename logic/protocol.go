package logic

type MessageType string

const (
	CONNECTED MessageType = "connected"
	NOTIFY    MessageType = "notify"
	CLOSED    MessageType = "closed"
)

type Message struct {
	Id      string
	Type    MessageType
	Payload string
}
