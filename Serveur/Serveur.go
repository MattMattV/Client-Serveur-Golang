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

func handleConnection(conn net.Conn) {

	var message string
	var decoder = gob.NewDecoder(conn)

	decoder.Decode(&message)
	fmt.Printf("Le client dit : %s\n", message)
}

func main() {

	ln, err := net.Listen("tcp", ":8080")
	verifErreur(err)
	fmt.Println("Socket d'écoute créé")

	for {
		
		conn, err := ln.Accept()
		verifErreur(err)
		
		fmt.Println("Connexion détectée")

		go handleConnection(conn)
	}
}