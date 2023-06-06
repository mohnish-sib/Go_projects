package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")//set the response type to json
	json.NewEncoder(w).Encode(movies)
}

func main(){
	r:=mux.NewRouter()
 
	movies =append(movies, Movie{ID:"1",Isbn: "444",Title: "Movie one",Director: &Director{Firstname:"John", Lastname:"Doe"}}) //& -> it gives address
	movies =append(movies, Movie{ID:"2",Isbn: "443",Title: "Movie two",Director: &Director{Firstname:"Mohnish", Lastname:"Lokhande"}}) 

	r.HandleFunc("/movies",getMovies).Methods("GET")
	// r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	// r.HandleFunc("/movies",createMovie).Methods("POST")
	// r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	// r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}