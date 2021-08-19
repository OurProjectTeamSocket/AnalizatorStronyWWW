package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func addwebsite(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil { // Parsowanie inputów
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	websitename := req.FormValue("addwebsite")
	username := req.FormValue("name")
	var masno string;

	// Usuwanie ostatniej litery z nazwy (masno) gdzie masno to jest nazwa usera ukrócona o 1 literę
	for nr, obj := range username { // usuwanie jednej litery z Go
		if nr < (len(username) - 1) {
			masno += string(obj);
		}
	}

	url := transforming(websitename)
	db, err := gorm.Open(sqlite.Open("db.db" ), &gorm.Config{})
	if err != nil {
		fmt.Println("Wyjebalo bleda spierdalaj")
	}

	var id_website User

	db.Table(masno + "_websites").Last(&id_website) // Branie ID

	db.Table(masno + "_websites").Create(&Websites{ID: id_website.ID+1, Url: websitename}) // Dodawanie strony do listy stron urzytkowanika

	db.Exec("CREATE TABLE " + masno + "_website_" + url + "(" + // Tworzenie DB da strony danego urzytkownika
		"ID INTEGER," +
		"Service INTEGER," +
		"Rang INTEGER," +
		"Days INTEGER," +
		"SSL INTEGER," +
		"Up INTEGER," +
		"Response DOUBLE," +
		"Down INTEGER" +
		")")

	db.Table(masno + "_website_" + url).Select("ID", "Service", "Rang", "Days", "SSL", "Up", "Response", "Down").Create(&Website{ID: 1, Service: true, Rang: 2, Days: 7, SSL: true, Up: 2, Response: 7.35312321, Down: 3}) // Podstawowe wartości przy dodawaniu strony

	result := db.Table("users").Where("name = ?", masno).Find(&User{}) // Id usera

	var id_user User

	result.First(&id_user)

	c := &http.Cookie {
		Name: "id",
		Value:  strconv.Itoa(id_user.ID),
		HttpOnly: true,
	}

	// Ustawianie odpowiadacza dla ciastka
	http.SetCookie(w, c)

	// Utworzenie ciastka (dodanie go)
	req.AddCookie(c)

	http.Redirect(w, req, "/dashboard", http.StatusSeeOther)

	//cos := beeep.Notify("Page status", "Page status is up", "");
}
