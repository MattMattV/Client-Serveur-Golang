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

func receiveMessage(conn net.Conn, isFinished chan bool) {

	var message string

	decoder := gob.NewDecoder(conn)

	for {
		
		decoder.Decode(&message)
		fmt.Println(message)

		if message == "BROADCAST" {
			isFinished <- true
		}
	}
}

func main() {

	var message string

	isFinished := make(chan bool)

	fmt.Println("Lancement du client")

	conn, err := net.Dial("tcp", "localhost:8080")
	checkError(err)
	
	encoder := gob.NewEncoder(conn)

	fmt.Println("Connexion réussie")

	fmt.Println("Entrez votre identifiant :")
	fmt.Scanln(&message)

	// envoi du nom
	checkError(encoder.Encode(message))
	fmt.Println("Le message à bien été envoyé.")

	go receiveMessage(conn, isFinished)
	
	<-isFinished

	fmt.Println("Le client va quitter...")
}
