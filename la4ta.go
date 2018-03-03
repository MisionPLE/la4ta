package main

import (
	"fmt"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	r.ParseForm()
	// logic part of log in
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
	fmt.Println(r.Form)
}

func main() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
