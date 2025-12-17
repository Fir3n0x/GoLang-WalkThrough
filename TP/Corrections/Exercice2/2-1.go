package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go '<hash>' <wordlist>")
        return
    }

    // Retrieve parameters
    target := os.Args[1]
    wordlist := os.Args[2]

    fmt.Println("Hash : ", target, ", longueur : ", len(target))

    // Open file
    file, err := os.Open(wordlist)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    // Read file
    var words []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    // Start timer
    start := time.Now()
    found := ""

    fmt.Println("Cassage du hash ", target, "...")

    // Break hash
    for _, word := range words {
        err = bcrypt.CompareHashAndPassword([]byte(target), []byte(strings.TrimSpace(word)))
        if err == nil {
            found = word
            break
        }
    }

    elapsed := time.Since(start)
    if found != "" {
        fmt.Printf("Mot de passe trouvé: %s\n", found)
    } else {
        fmt.Println("Mot de passe non trouvé.")
    }
    fmt.Printf("Temps: %s\n", elapsed)
}