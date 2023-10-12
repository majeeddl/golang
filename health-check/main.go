package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("Hello World")

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("HOST_NAME")
	port := os.Getenv("HOST_PORT")
	user := os.Getenv("HOST_USER")
	password := os.Getenv("HOST_PASSWORD")
	privateKey := os.Getenv("PRIVATE_KEY")

	fmt.Println(host, port, user)
	//fmt.Println(privateKey)
	// read private key file with passphrase
	// privateKey, err := os.ReadFile("~/Documents/ssh/majeed")

	if err != nil {
		fmt.Println("Error reading private key file")
	}

	fmt.Println(privateKey)

	signer, err := ssh.ParsePrivateKeyWithPassphrase([]byte(privateKey), []byte(password))

	if err != nil {
		fmt.Println("Error parsing private key file")
	}

	fmt.Println("Signer")
	fmt.Println(signer)

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
	}

	fmt.Println("Config")
	fmt.Println(config)

	var conn *ssh.Client

	conn, err = ssh.Dial("tcp", host+":"+port, config)

	if err != nil {
		fmt.Println("Error connecting to server")
	}

	defer conn.Close()
}
