package main

import (
    "bufio"
    "encoding/base64"
    "fmt"
    "net/http"
    "os"
    "sync"
)

func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for encoded := range jobs {
		// TODO : Décoder base64

		// TODO : Requête HTTP
		
		// TODO : Vérification du code retour pour filtrer les bonnes réponses
	}
}

func main() {
	// TODO : Ouvrir le fichier
	
	// TODO : Initialisation channels et synchronisation variable
	
	// TODO : Lecture du fichier
	
	// TODO : Collecte des résultats
	
	// TODO : Affichage des résultats
}
