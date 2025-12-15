package main

import(
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

	// Récupération des paramètres
	target := os.Args[1]
	wordlist := os.Args[2]

	fmt.Println("Hash : ", target, ", longueur : ", len(target))

	// TODO: Ouvrir le fichier
	
	// TODO : Lire le fichier
	
	// Démarrer le chronomètre
	start := time.Now()

	found := ""
	fmt.Println("Cassage du hash ", target, "...")

	// TODO : Casser le hash
	
	// TODO : Affichage du résultat
	elapsed := time.Since(start)
}
