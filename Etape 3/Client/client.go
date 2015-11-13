package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func checkError(err error) {
	if err != nil {
		log.Panicf("Une erreur est survenue : \n\t%v", err)
	}
}

func main() {

	var message string

	var iterations = 1

	fmt.Println("Lancement du client")

	conn, err := net.Dial("tcp", "localhost:8080")
	checkError(err)
	
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	fmt.Println("Connexion réussie")

	fmt.Println("Entrez le message à envoyer :")

	fmt.Scanln(&message)

	err = encoder.Encode(message)
	checkError(err)

	fmt.Println("Le message à bien été envoyé.")

	for {

		decoder.Decode(&message)
		fmt.Printf("[%d] Le serveur dit : %s\n", iterations, message)

		checkError(encoder.Encode(message))

		fmt.Printf("[%d] Envoyé : %s\n", iterations, message)

		iterations++
	}
}
