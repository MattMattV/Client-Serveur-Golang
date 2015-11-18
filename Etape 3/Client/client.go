package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
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
		
	decoder.Decode(&message)
	fmt.Printf("SERVEUR : %s\n", message)
}

func main() {

	fmt.Println("Lancement du client")

	source        := rand.NewSource(time.Now().UnixNano())
	randGenerator := rand.New(source)

	iterations := randGenerator.Intn(5000 - 250) + 250

	for i := 0; i < iterations; i++ {

		conn, err := net.Dial("tcp", "localhost:8080")
		checkError(err)
		
		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

		fmt.Println("Connexion réussie")

		//Génération d'un entier qui est transformé en string pour servir d'identifiant
		message := strconv.Itoa(randGenerator.Intn(123456))
		fmt.Printf("Je suis : %s\n", message)

		checkError(encoder.Encode(message))
		
		fmt.Println("Le message à bien été envoyé.")

		decoder.Decode(&message)
		fmt.Printf("SERVEUR : %s\n", message)
		
		
		//temporisation
		msSleepTime := randGenerator.Intn(10000 - 2000) - 2000
		time.Sleep(time.Duration(msSleepTime) * time.Millisecond)
		
		checkError(encoder.Encode("DISCONNECT"))
		conn.Close()

		fmt.Println("\n")
	}

	fmt.Printf("%d itérations\n", iterations)	
}
