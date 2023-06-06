package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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

func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)//we are taking all movies before the index and all movies after the index expect which we have to delete
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	id:=mux.Vars(r)["id"]

	movie:=Movie{}

	for _, item := range movies {
		if item.ID == id {
			movie = item
			break
		}
	}

	json.NewEncoder(w).Encode(movie) //this will return the response
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))

	movies =append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//set json content type
	w.Header().Set("Content-Type","application/json")

	// params
	params:=mux.Vars(r)

	//loop over the movies, range
	for index, item := range movies {
		if item.ID == params["id"]{
			movies =append(movies[:index],movies[index+1:]...)
			var movie Movie
			_ =json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	//delete the movie with the i.d that you have sent
	//add a new movie, with updated data using same id
	// not an ideal method, should not do in real projects
}

func main(){
	r:=mux.NewRouter()
 
	movies =append(movies, Movie{ID:"1",Isbn: "444",Title: "Movie one",Director: &Director{Firstname:"John", Lastname:"Doe"}}) //& -> it gives address
	movies =append(movies, Movie{ID:"2",Isbn: "443",Title: "Movie two",Director: &Director{Firstname:"Mohnish", Lastname:"Lokhande"}}) 

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at 8080\n")
	log.Fatal(http.ListenAndServe(":8080",r))
}