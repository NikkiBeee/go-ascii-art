package main

import (
		"fmt"
		"net/http"
		"html/template")


func indexHandler (w http.ResponseWriter, r *http.Request){
	t, err := template.ParseFiles("index.html")
	fmt.Println(err)
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)
}