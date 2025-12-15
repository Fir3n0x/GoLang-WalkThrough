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
		// Décodage Base64
		decoded, err := base64.StdEncoding.DecodeString(encoded)
		if err != nil {
			continue
		}
		url := "http://localhost:8080" + string(decoded)

		// Requête HTTP
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		resp.Body.Close()

		// Vérification du code
		if resp.StatusCode >= 200 && resp.StatusCode < 400 || resp.StatusCode == 403 || resp.StatusCode == 401 {
			results <- fmt.Sprintf("%s -> %d", url, resp.StatusCode)
		}
	}
}

func main() {
	file, err := os.Open("paths.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jobs := make(chan string, 100)
	results := make(chan string, 100)

	var wg sync.WaitGroup

	// Lancement de 10 workers
	for w := 1; w <= 10; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Lecture du fichier
	scanner := bufio.NewScanner(file)
	go func() {
		for scanner.Scan() {
			jobs <- scanner.Text()
		}
		close(jobs)
	}()

	// Collecte des résultats
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Valid path:", r)
	}
}
