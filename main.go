package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){ //any api route has two things request and response. here r is pointing to the request

	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w,"method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request){ 
	if err:=r.ParseForm(); err!=nil{
		fmt.Fprintf(w,"ParseForm() err: %v\n",err)
		return
	}
	fmt.Fprintf(w,"POST request successful\n")
	name:=r.FormValue("name")
	address:=r.FormValue("address")
	fmt.Fprintf(w,"Name =%s\n",name)
	fmt.Fprintf(w,"Address =%s\n",address)
}
func main() {

	fileServer := http.FileServer(http.Dir("./static")) //It will provide all file in the static folder to the server. we can open them loaclholst:8080/filename.html
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Staring server at port 8080\n")
	if err :=http.ListenAndServe(":8080",nil); err!=nil { //this starts the server
		log.Fatal(err)
	}
}
//https://youtu.be/jFfo23yIWac