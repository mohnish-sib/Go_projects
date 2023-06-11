package routes

import (
	"github.com/gorilla/mux"
	"github.com/mohnish-sib/Go_projects/pkg/controllers" // in go lang we always have absolute paths, unlike node.js
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{BookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("book/{bookId}",controllers.DeleteBook).Methods("DELETE")

}