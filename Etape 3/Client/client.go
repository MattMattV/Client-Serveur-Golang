package main

import (
	"encoding/gob"
	"fmt"
	"strings"
	"log"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Panicf("Une erreur est survenue : \n\t%v", err)
	}
}

func receiveMessage(conn net.Conn) {

	var message string

	decoder := gob.NewDecoder(conn)

	for {
		
		decoder.Decode(&message)
		fmt.Printf("SERVEUR : %s\n", message)

		if strings.Contains(message, "plein") {
			os.Exit(1)
		}
	}
}

func main() {

	signalCatcher := make(chan os.Signal, 1)
	signal.Notify(signalCatcher, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Lancement du client")

	conn, err := net.Dial("tcp", "localhost:8080")
	checkError(err)
	
	encoder := gob.NewEncoder(conn)

	fmt.Println("Connexion réussie")

	//fmt.Println("Entrez votre identifiant :")
	//fmt.Scanln(&message)

	//Génération d'un entier qui est transformé en string pour servir d'identifiant
	source        := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	message := strconv.Itoa(randGenerator.Int())

	fmt.Printf("Je suis : %s\n", message)

	// envoi du nom
	checkError(encoder.Encode(message))
	fmt.Println("Le message à bien été envoyé.")

	go receiveMessage(conn)
	
	if <-signalCatcher == syscall.SIGINT {
		encoder.Encode("DISCONNECT")
		fmt.Println("\n\nLe client va quitter...")
		return
	}

}
