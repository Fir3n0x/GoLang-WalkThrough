
package main

import (
    "fmt"
    "io"
    "net/http"
		"strconv"
)

// TODO : Créer les différents endpoints pour interagir avec le serveur

// endpoint
func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "")
}

func main() {
    http.HandleFunc("/", Handler)

    fmt.Println("Serveur démarré sur le port 8593...")
    http.ListenAndServe(":8593", nil)
}
