package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error al conectar con el servidor:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado al Servidor de Chat Empresarial.")

	// Goroutine para leer mensajes que envía el servidor
	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("\nConexión cerrada por el servidor.")
				os.Exit(0)
			}
			fmt.Print(msg)
		}
	}()

	// Leer desde la consola del usuario y enviar al servidor
	stdReader := bufio.NewReader(os.Stdin)
	for {
		text, _ := stdReader.ReadString('\n')
		text = strings.TrimSpace(text)

		_, err := conn.Write([]byte(text + "\n"))
		if err != nil {
			fmt.Println("Error al enviar mensaje:", err)
			break
		}

		if text == "SALIR" {
			break
		}
	}
}
