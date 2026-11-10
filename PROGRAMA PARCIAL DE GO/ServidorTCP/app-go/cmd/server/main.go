package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"parcialgo/internal/models"
	"parcialgo/internal/storage"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
		return
	}
	defer listener.Close()

	db := storage.NewChatMemory()
	fmt.Println("Servidor de Chat Empresarial iniciado en el puerto 8080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error al aceptar conexión:", err)
			continue
		}
		go handleConnection(conn, db)
	}
}

func handleConnection(conn net.Conn, db *storage.ChatMemory) {
	defer conn.Close()

	conn.Write([]byte("--- BIENVENIDO AL CHAT EMPRESARIAL ---\nIngresa tu usuario:\n"))

	reader := bufio.NewReader(conn)
	username, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	username = strings.TrimSpace(username)

	conn.Write([]byte(fmt.Sprintf("Hola %s. Escribe tus mensajes, 'HISTORIAL' para ver mensajes anteriores, o 'SALIR':\n", username)))

	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		text := strings.TrimSpace(input)

		if text == "SALIR" {
			conn.Write([]byte("Conexión cerrada. ¡Adiós!\n"))
			break
		}

		if text == "HISTORIAL" {
			history := db.GetHistory()
			conn.Write([]byte("--- HISTORIAL DE CONVERSACIONES ---\n"))
			if len(history) == 0 {
				conn.Write([]byte("(No hay mensajes previos)\n"))
			} else {
				for _, h := range history {
					conn.Write([]byte(h + "\n"))
				}
			}
			conn.Write([]byte("-----------------------------------\n"))
			continue
		}

		msg, err := models.NewMessage(username, text)
		if err != nil {
			conn.Write([]byte("Error de validación: " + err.Error() + "\n"))
			continue
		}

		formattedMsg := fmt.Sprintf("[%s]: %s", msg.Sender, msg.Content)
		db.SaveMessage(formattedMsg)

		conn.Write([]byte("Mensaje enviado y guardado correctamente.\n"))
	}
}
