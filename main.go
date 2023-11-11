package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/teerapoom/miniprojest_api002/control"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", control.GetMovies).Methods("GET")
	r.HandleFunc("/movise/{id}", control.GetMovie).Methods("GET")
	r.HandleFunc("/movise", control.CreateMovie).Methods("POST")
	r.HandleFunc("/movise/{id}", control.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movise/{id}", control.DeleteMovie).Methods("DELETE")
	fmt.Println("Starting server at port 3030")
	log.Fatal(http.ListenAndServe(":3030", r)) //เริ่ม HTTP server บนพอร์ต 8000 ถ้า Error -> exit
}
