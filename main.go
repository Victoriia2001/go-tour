package main

import ( 
	"net/http"
	"log"
	"html/template"
)

func main() {
	http.HandleFunc("/", mainPage)
	port := ":9999"
	println("Server listen on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func mainPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("static/ind.html")
	if err != nil {
	   http.Error(w, err.Error(), 400)
	   return	
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	   
}