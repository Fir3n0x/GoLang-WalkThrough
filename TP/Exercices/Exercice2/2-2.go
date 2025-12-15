package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"golang.org/x/crypto/bcrypt"
)

func worker(words []string, target string, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	for _, word := range words {
		// TODO : Comparer le hash avec le mot de passe candidat
	}
}


func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go '<hash>' <wordlist>")
		return
	}

	// Récupération des paramètres
	target := os.Args[1]
	wordlist := os.Args[2]

	fmt.Println("Hash : ", target, ", longueur : ", len(target))

	// TODO : Ouvrir le fichier
	
	// TODO : Lecture du fichier

	// Démarrer le chronomètre
	start := time.Now()
	
	// Variables pour channel et synchronisation
	result := make(chan string, 1)
	var wg sync.WaitGroup

	// TODO : Assigner un espace de recherche pour chaque worker
	chunkSize := 100
	for i := 0; i<len(); i += chunkSize {

	}

	// TODO : Go routine pour attendre que tous les workers aient fini
	
	// TODO : Regarder la réponse du channel
	// Si le hash est trouvé alors break
	found := ""
	
	// TODO : Affichage du résultat
	elapsed := time.Since(start)
}
