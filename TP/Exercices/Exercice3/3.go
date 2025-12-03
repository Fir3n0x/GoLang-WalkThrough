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

        conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
        if err != nil {
            return err
        }

        client := &Client{Conn: conn, Send: make(chan []byte)}
        mutex.Lock()
        clients[client] = true
        mutex.Unlock()

        go handleRead(client)
        go handleWrite(client)

        return nil
    })

    // Goroutine to broadcast messages
    go func() {
        // TODO
    }()

    e.Logger.Fatal(e.Start(":8080"))
}

func handleRead(client *Client) {
    // TODO
}

func handleWrite(client *Client) {
    // TODO
}