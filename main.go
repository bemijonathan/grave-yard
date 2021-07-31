package main

import (
	"fmt"
	"net/http"

	"github.com/bemijonathan/grave-yard/controllers"
	"github.com/bemijonathan/grave-yard/db"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home).Methods(http.MethodGet)
	r.HandleFunc("/projects", controllers.Allpost).Methods(http.MethodGet)

	fileServer := http.FileServer(http.Dir("./assets/"))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	port := 8080

	db.ConnectToDatabase()

	db.ServerStarted()

	fmt.Println("server is running on ", port)
	http.ListenAndServe(":8080", r)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello from  hrllo")
}
