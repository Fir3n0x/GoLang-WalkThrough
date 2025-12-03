package main

import (
    "fmt"
    "net/http"
    "sync"

    "github.com/gorilla/websocket"
    "github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
    Conn *websocket.Conn
    Send chan []byte
	Username string
}

var (
    clients   = make(map[*Client]bool)
    broadcast = make(chan []byte)
    mutex     sync.Mutex
)

func main() {
    e := echo.New()

    // WebSocket route
    e.GET("/ws", func(c echo.Context) error {

		username := c.QueryParam("username")
		if username == "" {
			username = "Anonymous"
		}

        conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
        if err != nil {
            return err
        }

        client := &Client{Conn: conn, Send: make(chan []byte), Username: username}
        mutex.Lock()
        clients[client] = true
        mutex.Unlock()

        go handleRead(client)
        go handleWrite(client)

        return nil
    })

    // Goroutine to broadcast messages
    go func() {
        for {
            msg := <-broadcast
            mutex.Lock()
            for client := range clients {
                select {
                case client.Send <-  msg:
                default:
                    close(client.Send)
                    delete(clients, client)
                }
            }
            mutex.Unlock()
        }
    }()

    e.Logger.Fatal(e.Start(":8080"))
}

func handleRead(client *Client) {
    defer func() {
        mutex.Lock()
        delete(clients, client)
        mutex.Unlock()
        client.Conn.Close()
    }()

    for {
        _, msg, err := client.Conn.ReadMessage()
        if err != nil {
            break
        }
        fmt.Printf("Message reçu: %s~: %s\n", client.Username, string(msg))
		fullMsg := fmt.Sprintf("%s~: %s", client.Username, string(msg))
        broadcast <- []byte(fullMsg)
    }
}

func handleWrite(client *Client) {
    for msg := range client.Send {
        err := client.Conn.WriteMessage(websocket.TextMessage, msg)
        if err != nil {
            break
        }
    }
}