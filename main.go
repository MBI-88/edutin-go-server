package main

import (
	"fmt"
	"log"
	"net/http"
	"web/server"
)

func main() {
	fmt.Println("[*] Server running in port 80")
	router := server.Router()
	if err := http.ListenAndServe(":80", router); err != nil {
		log.Fatalf("%s\n", err)
	}
}
