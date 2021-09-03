package main

import (
	"fmt"
	adds "github.com/XploraTechProducts/monitorAI/controllers"
	"net/http"
)

// Główny serwer

func main() {
	http.HandleFunc("/login", adds.Login())
	http.HandleFunc("/register", adds.Register())
	http.HandleFunc("/dashboard", adds.Dashboard()) // dashboard-analytics.html
	http.HandleFunc("/addwebsite", adds.Addwebsite())
	http.HandleFunc("/logout", adds.Logout())
	http.Handle("/", http.FileServer(http.Dir("./Website/")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Nie udało się uruchomić")
	}
}