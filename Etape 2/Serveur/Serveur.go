package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func checkError(err error) {
	if err != nil {
		log.Panicf("Une erreur est survenue : \n\t%v", err)
	}
}


func handleConnection(conn net.Conn) {

	var message string

	var iterations = 1
	var decoder    = gob.NewDecoder(conn)
	var encoder    = gob.NewEncoder(conn)

	for {
		
		checkError(decoder.Decode(&message))
		fmt.Printf("[%d] Le client dit : %s\n", iterations, message)

		checkError(encoder.Encode(message))
		fmt.Printf("[%d] Message envoyé %s\n",iterations, message)

		iterations++
	}
}

func main() {

	maxClients, err := strconv.Atoi(os.Getenv("MAX_CLIENTS"))
	checkError(err)

	fmt.Printf("Nombre maximal de connexions : %d", maxClients)


	ln, err := net.Listen("tcp", ":8080")
	checkError(err)
	fmt.Println("Socket d'écoute créé")

	defer fmt.Println("Le serveur est fermé.")
	defer ln.Close()

	for {
		
		conn, err := ln.Accept()
		checkError(err)
		
		fmt.Println("Connexion détectée")

		go handleConnection(conn)
	}
}