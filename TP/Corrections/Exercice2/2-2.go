package main

import(
	"fmt"
	"time"
	"golang.org/x/crypto/bcrypt"
	"os"
	"sync"
	"bufio"
)

// Worker function
func worker(words []string, target string, wg *sync.WaitGroup, result chan<- string) {
	defer wg.Done()
	for _, word := range words {
		// Compare hash with password candidate
		if bcrypt.CompareHashAndPassword([]byte(target), []byte(word)) == nil {
			result <- word
			return
		}
	}
}

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
	// Channel variable
	result := make(chan string, 1)
	// Wait for all workers to finish (variable)
	var wg sync.WaitGroup

	chunkSize := 100 // tradeoff
	// Assign research space for each worker
	for i := 0; i < len(words); i += chunkSize {
		end := i + chunkSize
		if end > len(words) {
			end = len(words)
		}
		// Add worker to wg variable
		wg.Add(1)
		// Launch go routine (in parallel)
		go worker(words[i:end], target, &wg, result)
	}

	// Go routine to wait for all workers to finish
	go func() {
		wg.Wait()
		close(result)
	}()

	// Check the channel's response
	// If hash found then break
	found := ""
	for r := range result {
		found = r
		break
	}

	elapsed := time.Since(start)
	if found != "" {
		fmt.Printf("Mot de passe trouvé: %s\n", found)
	}else{
		fmt.Println("Mot de passe non trouvé.")
	}
	fmt.Printf("Temps: %s\n", elapsed)

}