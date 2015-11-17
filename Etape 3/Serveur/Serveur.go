package main

import (
	"encoding/gob"
	"fmt"
	"io"
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
	
	if err != nil && err != io.EOF {
		log.Panicf("Une erreur est survenue : \n\t%v", err)
	}
}

func handleConnection(conn net.Conn, chName chan string, chStatus chan net.Conn) {

	var message string

	//ces deux variables vont permttre de dialoguer avec le client
	var decoder = gob.NewDecoder(conn)
	

	checkError(decoder.Decode(&message))
	fmt.Printf("New client : %s\n", message)

	// on envoie le message dans le main pour que le tableau
	// de Client soit complété
	chName <- message

	sendMessage("Vous êtes bien connecté", conn)

	//on récupere le signal de déconnexion pour que le main puisse supprimer le client
	for {

		if message == "DISCONNECT" {
			chStatus <- conn
		} else {
			chStatus <- nil
		}
		
		checkError(decoder.Decode(&message))
	}
}

func broadcast(message string, mapClient map[string]net.Conn,) {

	for key, value := range mapClient {

		sendMessage("BROADCAST !!\n\t" + message, value)
		fmt.Printf("\tEnvoi vers %s\n", key)
	}
}

func sendMessage(message string, conn net.Conn) {

	encoder := gob.NewEncoder(conn)

	checkError(encoder.Encode(message))
}

func main() {

	//variables
	var nbConnected = 0

	var chName      = make(chan string)
	var chStatus    = make(chan net.Conn)
	var mapClient   = make(map[string]net.Conn)

	//on vérifie la présence de la variable d'environnement
	maxClients, err := strconv.Atoi(os.Getenv("MAX_CLIENTS"))
	checkError(err)

	fmt.Printf("\nMaximal connections : %d\n\n", maxClients)
	
	//le serveur écoute sur le port 8080 avec le protocole TCP
	ln, err := net.Listen("tcp", ":8080")
	checkError(err)

	fmt.Println("Listen socket created\n")

	for {
		
		fmt.Println("debut attente")
		conn, err := ln.Accept()
		checkError(err)
		fmt.Println("fin attente")

		
		//on ajoute un client seulement si il y a de place
		if nbConnected < maxClients {
			
			//on crée une goroutine pour accepter plusieurs clients
			go handleConnection(conn, chName, chStatus) 
			
			//la gouroutine assiciée à un client va envoyer l'identifiant pour que l'on puisse l'enregistrer dans une map
			id := <-chName
			mapClient[id] = conn
			nbConnected++

		} else {

			//on envoie un message au client pour qu'il sache que le serveur est plein
			sendMessage("Impossible de se connecter, serveur plein", conn)
			
			fmt.Println("debut attente")
			conn, err = ln.Accept()
			checkError(err)
			fmt.Println("fin attente")
		}
	}

	fmt.Printf("1 : nbConnected %d/%d\n", nbConnected, maxClients)
	//on recupère un signal de déconnexion pour faire de la place pour un autre client
	var status = <-chStatus
	fmt.Println("DEBUG | status", status)
	if status != nil {
		
		//on recherche l'Id associé à l'objet net.Conn que la goroutine à envoyé dans status
		for key, value := range mapClient {

			fmt.Printf("DEBUG | key: %s, value: %v\n", key, value)

			if(value == status) {

				fmt.Printf("Suppression du client %s... ", key)
				delete(mapClient, key)
				fmt.Printf("TERMINE\n\n")
				nbConnected 
			}
	}

	//on dit à tout le monde que le serveur est plein
	if nbConnected == maxClients {
		fmt.Println("Server full !")
		broadcast("Server is full !", mapClient)
	}

		fmt.Println("affich encore")
		for key, value := range mapClient {
			
			fmt.Printf("DEBUG | key: %s, value: %v\n", key, value)
		}
	}
	
	fmt.Printf("2 : nbConnected %d/%d\n\n", nbConnected, maxClients)
}