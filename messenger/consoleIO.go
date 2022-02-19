package messenger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MessengerConsole struct{}

func (io MessengerConsole) SendMessage(message string) {
	fmt.Println(message)
}

func (io MessengerConsole) SendMultiline(message MultilineMessage) {
	for _, s := range message {
		fmt.Println(s)
	}
}

func (io MessengerConsole) ReadData() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	data, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return data, nil
}
