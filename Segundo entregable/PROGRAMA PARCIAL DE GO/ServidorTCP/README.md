# Prototipo de Chat Empresarial TCP en Go

Este proyecto implementa una arquitectura cliente-servidor basada en el protocolo TCP utilizando Go, desarrollada como entregable para la Semana 10. El sistema garantiza la modularidad, concurrencia y validación estricta de mensajes en memoria de forma segura.

## 📁 Estructura del Proyecto

El proyecto sigue estrictamente el estándar de organización para aplicaciones en Go:

```text
ServidorTCP/
├── app-go/
│   └── cmd/
│       ├── server/
│       │   └── main.go      # Lógica principal del Servidor TCP (Puerto 8080)
│       └── client/
│           └── main.go      # Lógica del Cliente interactivo
├── internal/
│   ├── models/
│   │   └── chat.go          # Validación de estructuras y mensajes
│   └── storage/
│       └── memory.go        # Almacenamiento seguro en memoria con sync.Mutex
└── go.mod                   # Definición del módulo (parcialgo)


## 🎥 Video de Demostración
Puedes ver la ejecución y explicación del proyecto en el siguiente enlace:
https://drive.google.com/file/d/1tTpz_YyDOkplSZrAoohd11BtmVi4joNY/view?usp=sharing