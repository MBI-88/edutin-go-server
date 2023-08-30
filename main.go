package main

import (
	"fmt"
	"log"
	"net/http"
	"web/server"
)

func main() {
	fmt.Println("[*] Server running in localhost port 8000")
	router := server.Router()
	if err := http.ListenAndServe("localhost:8000", router); err != nil {
		log.Fatalf("%s\n", err)
	}

	

}
