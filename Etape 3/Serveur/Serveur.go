package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
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

func handleConnection(conn net.Conn, mapClient map[string]net.Conn) {

	var message, id string

	// le mutex va servir à empêcher les autres goroutines d'accéder à la map
	// de la même façon que des sémaphores en C
	var mutex = &sync.Mutex{}

	// pour pouvoir recevoir les messages du client
	var decoder = gob.NewDecoder(conn)
	
	//réception du nom du client
	checkError(decoder.Decode(&message))
	
	id = message

	fmt.Printf("GOROUTINE : New client : %s\n", id)

	mutex.Lock()
	mapClient[id] = conn
	mutex.Unlock()

	sendMessage("GOROUTINE : Vous êtes bien connecté", conn)

	// on récupere le signal de déconnexion pour que le main puisse supprimer le client
	for {
		checkError(decoder.Decode(&message))

		if message == "DISCONNECT" {
			
			mutex.Lock()
			fmt.Println("AVANT\n")
			for key, value := range mapClient {
				fmt.Printf("[%s] -> %v\n", key, value)
			}
			mutex.Unlock()

			mutex.Lock()
			delete(mapClient, id)
			fmt.Printf("Client %s déconnecté\n", id)
			mutex.Unlock()

			mutex.Lock()
			fmt.Println("APRES")
			for key, value := range mapClient {
				fmt.Printf("[%s] -> %v\n", key, value)
			}
			fmt.Println()
			mutex.Unlock()

			return
		}		
	}
}

func broadcast(message string, mapClient map[string]net.Conn) {

	var mutex = sync.Mutex{}

	mutex.Lock()
	for key, value := range mapClient {

		sendMessage("BROADCAST !!\n\t" + message, value)
		fmt.Printf("\tEnvoi vers %s\n", key)
	}

	mutex.Unlock()
}

func sendMessage(message string, conn net.Conn) {

	encoder := gob.NewEncoder(conn)

	checkError(encoder.Encode(message))
}

func main() {

	// variable
	var mapClient   = make(map[string]net.Conn)

	// on vérifie la présence de la variable d'environnement
	maxClients, err := strconv.Atoi(os.Getenv("MAX_CLIENTS"))
	checkError(err)

	fmt.Printf("\nMaximal connections : %d\n\n", maxClients)
	
	// le serveur écoute sur le port 8080 avec le protocole TCP
	ln, err := net.Listen("tcp", ":8080")
	checkError(err)

	fmt.Println("Listen socket created\n")

	for {
		
		conn, err := ln.Accept()
		checkError(err)

		// on ajoute un client seulement si il y a de place
		if len(mapClient) < maxClients {
			
			// on crée une goroutine pour accepter plusieurs clients en même temps
			go handleConnection(conn, mapClient) 
			
		} else {

			// on envoie un message au client pour qu'il sache que le serveur est plein
			sendMessage("Impossible de se connecter, serveur plein", conn)
		}

		fmt.Printf("capacité : %d/%d\n", len(mapClient), maxClients)
		
		// on dit à tout le monde que le serveur est plein
		if len(mapClient) == maxClients {
			fmt.Println("Server full !")
			broadcast("Server is full !", mapClient)
		}

		//time.Sleep(time.Second)		
	}
}