package main

import (
	"calculator/rps"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func calcutator(w http.ResponseWriter, r *http.Request) {
	//num1, _ := strconv.ParseFloat(r.PostFormValue("number1"), 64)
	//num1 := document.getElementById("number1").value
	//num1, _ := strconv.ParseFloat("nm1", 64)
	//num1 := 10.0
	num2 := 10.0
	//num1, _ := strconv.ParseFloat(r.URL.Query().Get("n1"), 64)
	num1, _ := strconv.ParseFloat(r.FormValue("number1"), 64)
	operator, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.Calculation(num1, num2, operator)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func main() {
	http.HandleFunc("/calculate", calcutator)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
