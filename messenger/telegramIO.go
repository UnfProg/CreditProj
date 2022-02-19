package messenger

type MessengerTelegram struct{}

func (io MessengerTelegram) SendMessage(message string) {
	// to-do
}

func (io MessengerTelegram) SendMultiline(message MultilineMessage) {
	// to-do
}

func (io MessengerTelegram) ReadData() (int, error) {
	// to-do
	return 0, nil
}
