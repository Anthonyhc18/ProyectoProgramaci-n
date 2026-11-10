package storage

import (
	"sync"
)

// ChatMemory gestiona el almacenamiento temporal de mensajes en memoria de forma segura
type ChatMemory struct {
	mu       sync.Mutex
	messages []string
}

// NewChatMemory función constructora para la memoria
func NewChatMemory() *ChatMemory {
	return &ChatMemory{
		messages: make([]string, 0),
	}
}

// SaveMessage guarda un mensaje usando un receptor por puntero
func (m *ChatMemory) SaveMessage(msg string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.messages = append(m.messages, msg)
}

// GetHistory retorna una copia del historial para mantener el encapsulamiento
func (m *ChatMemory) GetHistory() []string {
	m.mu.Lock()
	defer m.mu.Unlock()

	historyCopy := make([]string, len(m.messages))
	copy(historyCopy, m.messages)
	return historyCopy
}
