package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
)

func verifErreur(err error) {
	if err != nil {
		log.Panicf("Une erreur est survenue : \n\t%v", err)
	}
}

func main() {

	var message = "Coucou !"

	fmt.Println("Lancement du client")

	conn, err := net.Dial("tcp", "localhost:8080")
	verifErreur(err)

	fmt.Println("Connexion réussie")

	encoder := gob.NewEncoder(conn)

	err = encoder.Encode(message)
	verifErreur(err)

	fmt.Println("Le message à bien été envoyé.")
}
