package messenger

type MultilineMessage []string

type Messenger interface {
	SendMessage(message string)
	SendMultiline(message MultilineMessage)
	ReadData() (int, error)
}
