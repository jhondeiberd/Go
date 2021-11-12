package models

import (
	"Go/Final_Project/src/config"
	"Go/Final_Project/src/entities"
)

type ClientModel struct {
}

func (*ClientModel) findall() ([]entities.Client, error) {
	db, err := config.GetDB()
	if err != nil {
		return nil, err
	} else {
		rows, err2 := db.Query("select * from client")
		if err2 != nil {
			return nil, err2
		} else {
			var clients []entities.Client
			for rows.Next() {
				var client []entities.Client
				rows.Scan(&client.idClient, &client.firstName, &client.lastName, &client.email)
				clients = append(clients, client)
			}
			return clients, nil
		}
	}
}
