package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func style(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "style.css")
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		log.Println(t.Execute(w, nil))
	} else {
		err := req.ParseForm()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("username: ", req.Form["username"])
		fmt.Println("password: ", req.Form["password"])

	}
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/style.css", style)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
