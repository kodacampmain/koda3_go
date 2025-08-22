package minitask3

import "log"

type Message struct {
	Sender  string
	Content string
}

func NewMessage(sender, content string) Message {
	return Message{
		Sender:  sender,
		Content: content,
	}
}
func (m Message) Send(chn chan Message) {
	chn <- m
}

type WhiteBoard struct {
	// messages []Message
}

func NewWhiteBoard() WhiteBoard {
	return WhiteBoard{}
}
func (w WhiteBoard) Listen(chn chan Message) {
	for msg := range chn {
		log.Printf("Sender: %s, Content: %s", msg.Sender, msg.Content)
	}
}
