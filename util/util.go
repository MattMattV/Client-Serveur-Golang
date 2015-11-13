package util

import (
	"encoding/gob"
	"log"
	"net"
)

func CheckError(err error) {
	if err != nil {
		log.Panicf("Connection error %v", err)
	}
}

func Send(message interface{}, conn net.Conn) error {
	return gob.NewEncoder(conn).Encode(&message)
}

func Receive(conn net.Conn) (message interface{}) {
	return gob.NewDecoder(conn).Decode(&message)
}
