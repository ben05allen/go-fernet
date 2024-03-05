package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fernet/fernet-go"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	keyString := os.Getenv("FERNET_KEY")
	token := os.Getenv("TOKEN")

	key, err := fernet.DecodeKey(keyString)
	if err != nil {
		log.Fatalf("Invalid key: %v", err)
	}

	message := fernet.VerifyAndDecrypt([]byte(token), 0, []*fernet.Key{key})
	if message == nil {
		log.Fatal("Invalid token or key")
	}

	fmt.Println(string(message))

}
