

package main

import (
    "fmt"
    "io"
    "net/http"
		"strconv"
)

var lastA int
var lastB int

// GET /ping -> "pong" ou le token
func pingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "pong")
}

// POST /flag1 -> reçoit fragment et répond "ok"
func flag1Handler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()
    fmt.Println("Fragment flag1 reçu:", string(body))
    fmt.Fprint(w, "ok")
}

// GET /echo?word=XYZ -> renvoie XYZ
func echoHandler(w http.ResponseWriter, r *http.Request) {
    word := r.URL.Query().Get("token")
    fmt.Fprint(w, word)
}

// GET /clientA?num=INT -> renvoie INT
func clientAHandler(w http.ResponseWriter, r *http.Request) {
    numStr := r.URL.Query().Get("num")
		n, _ := strconv.Atoi(numStr)
		lastA = n
    fmt.Fprint(w, numStr)
}

// GET /clientB?num=INT -> renvoie INT
func clientBHandler(w http.ResponseWriter, r *http.Request) {
    numStr := r.URL.Query().Get("num")
		n, _ :=strconv.Atoi(numStr)
		lastB = n
		fmt.Fprint(w, numStr)
}

// GET /result -> renvoie la somme
func resultHandler(w http.ResponseWriter, r *http.Request) {
    sum := lastA + lastB
    fmt.Fprint(w, sum)
}

// POST /flag2 -> reçoit fragment et répond "ok"
func flag2Handler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()
    fmt.Println("Fragment flag2 reçu:", string(body))
    fmt.Fprint(w, "ok")
}

// POST /flag3 -> reçoit fragment et répond "ok"
func flag3Handler(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    defer r.Body.Close()
    fmt.Println("Fragment flag3 reçu:", string(body))
    fmt.Fprint(w, "ok")
}

func main() {
    http.HandleFunc("/ping", pingHandler)
    http.HandleFunc("/flag1", flag1Handler)
    http.HandleFunc("/echo", echoHandler)
		http.HandleFunc("/clientA", clientAHandler)
		http.HandleFunc("/clientB", clientBHandler)
    http.HandleFunc("/result", resultHandler)
    http.HandleFunc("/flag2", flag2Handler)
    http.HandleFunc("/flag3", flag3Handler)

    fmt.Println("Serveur démarré sur le port 8593...")
    http.ListenAndServe(":8593", nil)
}
