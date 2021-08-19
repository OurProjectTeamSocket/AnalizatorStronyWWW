package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func login(w http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil { // Parsowanie inputów
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// inputy
	name := req.FormValue("user-name")
	password := req.FormValue("user-password")

	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(w, "Cannot connect to database, sorry form problems. Please contact with administrator of this website")
	}

	result := db.Table("users").Where("name = ?", name).Find(&User{}) // w TABELI users GDZIE name równa się zawartość zmiennej name

	var item User

	result.First(&item) // pierwszy item który spełnia wymagania ( name = ? )

	if !CheckHash(password, item.Password) { // Jeżeli hasła NIE są spójne
		fmt.Println("Bad password")
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	db = nil

	// Zmienna z adresem ID danego użytkownika
	id := item.ID

	// Tworzenie ciasteczka z danymi
	c := &http.Cookie {
		Name: "id",
		Value:  strconv.Itoa(id),
		HttpOnly: true,
	}

	// Ustawianie odpowiadacza dla ciastka
	http.SetCookie(w, c)

	// Utworzenie ciastka (dodanie go)
	req.AddCookie(c)

	// Przekierowywanie
	http.Redirect(w, req, "/dashboard", http.StatusSeeOther)

}
