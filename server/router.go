package server

import (
	"net/http"
)

// Router function
func Router() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/",Index)
	router.HandleFunc("/man",Man)
	router.HandleFunc("/women",Women)
	router.HandleFunc("/children",Children)
	router.HandleFunc("/offerts",Offerts)
	router.HandleFunc("/news",News)
	router.HandleFunc("/pays",Pays)
	router.HandleFunc("/car",Car)
	router.HandleFunc("/contact",Contact)
	router.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	return router
}