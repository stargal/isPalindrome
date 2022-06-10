package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("style.css")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	io.Copy(w, file)
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if req.Method == "GET" {
		t, _ := template.ParseFiles("login.html", "style.css")
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
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}

}
