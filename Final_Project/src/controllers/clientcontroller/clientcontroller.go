package clientcontroller

import (
	"Go/Final_Project/src/models"
	"net/http"
	"text/template"
)

func Index(response http.ResponseWriter, request *http.Request) {
	var clientModel models.ClientModel
	clients, _ := clientModel.FindAll()
	data := map[string]interface{}{
		"clients": clients,
	}
	tmp, _ := template.ParseFiles("views/client/index.html")
	tmp.Execute(response, data)
}
