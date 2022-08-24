package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/cafaray/internal/data"
	"github.com/cafaray/internal/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	// Connection to the database
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	// start the server
	go serv.Start()

	// wait for an interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attemp a graceful shutdown
	serv.Close()
	data.Close()
}

/*

var store = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

func main() {
	fmt.Println("Starting service ...")
}

func Put(key string, value string) error {
	store[key] = value
	// As this operation will be idempotent, we set the same value to returns
	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]
	if !ok {
		return "", ErrorNoSuchKey // <- here the previous definition of the error is key
		// to return an error defined when an element is not found
		// this kind is defined as a `sentinel error``
		// otherwise it returns an error

	}
	// The get operation is not idempotent, also is classified as nildepotemp
	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}
*/
