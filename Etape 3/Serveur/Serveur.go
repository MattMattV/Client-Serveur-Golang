package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

type Client struct {
	Id   string
	Conn net.Conn
}

func checkError(err error) {
	if err != nil {
		log.Panicf("Une erreur est survenue : \n\t%v", err)
	}
}


func handleConnection(conn net.Conn, channel chan string, estRempli bool) {

	var decoder    = gob.NewDecoder(conn)
	var encoder    = gob.NewEncoder(conn)
	var message    string

	checkError(decoder.Decode(&message))
	fmt.Printf("Nouveau client : %s\n", message)

	// on envoie le message "dans le main" pour que le tableau
	// de Client soit complété
	channel <- message

	checkError(encoder.Encode("Vous êtes bien connecté"))
}

func broadcast(message string, tab []Client) {
	
	var encoder *gob.Encoder

	for _, c := range tab {
		encoder = gob.NewEncoder(c.Conn)
		checkError(encoder.Encode("BROADCAST"))
		checkError(encoder.Encode(message))
		fmt.Printf("\tEnvoi vers %s\n", c.Id)
	}
}

func main() {

	var nbConnected = 0
	var channel = make(chan string)

	maxClients, err := strconv.Atoi(os.Getenv("MAX_CLIENTS"))
	checkError(err)

	fmt.Printf("\nNombre maximal de connexions : %d\n\n", maxClients)
	
	var tabClients = make([]Client, maxClients)

	ln, err := net.Listen("tcp", ":8080")
	checkError(err)
	fmt.Println("Socket d'écoute créé")

	defer fmt.Println("Le serveur est fermé.")
	defer ln.Close()

	for {
		
		if nbConnected == maxClients {
			fmt.Println("Serveur plein !")
			broadcast("Le serveur est rempli !", tabClients)
			return
		}

		conn, err := ln.Accept()
		checkError(err)
		
		fmt.Println("Connexion détectée")

		isFull := nbConnected == maxClients
		go handleConnection(conn, channel, isFull)

		if nbConnected < maxClients {
			
			tabClients[nbConnected] = Client{<-channel, conn}

		}

		nbConnected++
	}
}