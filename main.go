package main

import (
	"fmt"
	"net/http"
)

// Szablon do bazy danych

type User struct {
	ID int
	Name string
	Email string
	Password string
}

type Websites struct { // DB z listą stron usera
	ID int
	Url string
}

type Website struct { // DB z informacjami strony
	ID int
	Service bool
	Rang int
	Days int
	SSL bool
	Up int
	Response float64
	Down int
}

// Główny serwer

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/dashboard", dashboard) // dashboard-analytics.html
	http.HandleFunc("/addwebsite", addwebsite)
	http.Handle("/", http.FileServer(http.Dir("C:/Users/User/Desktop/monitorAI-main/")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Nie udało się uruchomić")
	}
}