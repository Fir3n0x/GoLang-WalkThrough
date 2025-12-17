package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	hashed, err := bcrypt.GenerateFromPassword([]byte(""), bcrypt.DefaultCost)
	if err != nil {
        panic(err)
    }
	fmt.Println(string(hashed), len(hashed))
}
