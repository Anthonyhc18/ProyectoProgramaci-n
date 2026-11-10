package models

import "errors"

type Message struct {
	Sender  string
	Content string
}

func NewMessage(sender, content string) (*Message, error) {
	if sender == "" {
		return nil, errors.New("el remitente no puede estar vacío")
	}
	if content == "" {
		return nil, errors.New("el contenido del mensaje no puede estar vacío")
	}
	return &Message{
		Sender:  sender,
		Content: content,
	}, nil
}
