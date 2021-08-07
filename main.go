package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

// Szablon do bazy danych

type User struct {
	ID int
	Name string
	Email string
	Password string
}

// Zabezpieczenia

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Łączenie z serwerem

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
		fmt.Fprintf(w, "Can not connect to database, sorry form problems. Please contact with administrator of this website")
	}

	result := db.Table("users").Where("name = ?", name).Find(&User{})

	var item User

	result.First(&item) // pierwszy item który spełnia wymagania ( name = ? )

	if !CheckHash(password, item.Password) { // Jeżeli hasła NIE są spójne
		fmt.Println("Bad password")
		return
	}

	fmt.Fprintf(w, "Logged!")

	db = nil

	http.Redirect(w, req, "https://www.google.pl", http.StatusSeeOther) // Przeniesienie na goole.pl

}

func register(w http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := req.FormValue("user-name")
	email := req.FormValue("user-email")
	password := req.FormValue("user-password")

	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(w, "Can not connect to database, sorry form problems. Please contact with administrator of this website")
	}

	result := db.Table("users").Where("name = ?", name).Find(&User{})

	if result.RowsAffected > 0 { // Jeżeli pojawiły się wyniki wyszukiwania
		fmt.Fprintf(w, "Account exist")
		return
	}

	var item User

	result = db.Table("users").Last(&item) // pobieranie odstaniniego rekordu ( najwyższego ID )

	pass, _ := Hash(password) // Haszowanie hasła

	db.Table("users").Select("ID", "Name", "Email", "Password").Create(&User{ID: item.ID+1 /* dodawanie 1 do najwyższego ID w bezie danych */, Name: name, Email: email, Password: pass})

	db = nil

}

func monitor(w http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(w, "Just Monitor :)") // Or Monika ;)
}

// Główny serwer

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/monitor", monitor)
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Nie udało się uruchomić")
	}
}