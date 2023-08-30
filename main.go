package main

import (
	"fmt"
	"log"
	"net/http"
	"web/server"
)

func main() {
	fmt.Println("[*] Server running in localhost port 8080")
	router := server.Router()
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("%s\n", err)
	}

	

}
