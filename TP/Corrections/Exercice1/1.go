package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Ouvrir le fichier
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Erreur ouverture fichier:", err)
        return
    }
    defer file.Close()

    // Lire les nombres
    var numbers []int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        n, err := strconv.Atoi(scanner.Text())
        if err == nil {
            numbers = append(numbers, n)
        }
    }

    // Channel pour les résultats
    results := make(chan int)

    // Lancer les goroutines
    for _, num := range numbers {
        go func(n int) {
            results <- n * n // calcul du carré
        }(num)
    }

    // Récupérer les résultats
    for i := 0; i < len(numbers); i++ {
        fmt.Println(<-results)
    }
}