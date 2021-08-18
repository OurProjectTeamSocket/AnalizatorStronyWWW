package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Szablon do bazy danych

type User struct {
	ID int
	Name string
	Email string
	Password string
}

type Websites struct {
	ID int
	Url string
}

type Website struct {
	ID int
	Service bool
	Rang int
	Days int
	SSL bool
	Up int
	Response float64
	Down int
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
		fmt.Fprintf(w, "Cannot connect to database, sorry form problems. Please contact with administrator of this website")
	}

	result := db.Table("users").Where("name = ?", name).Find(&User{})

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

	result = db.Table("users").Last(&item) // pobieranie ostatniego rekordu ( najwyższego ID )

	pass, _ := Hash(password) // Haszowanie hasła
	db.Table("users").Select("ID", "Name", "Email", "Password").Create(&User{ID: item.ID+1 /* dodawanie 1 do najwyższego ID w bezie danych */, Name: name, Email: email, Password: pass})

	db = nil

	http.Redirect(w, req, "/", http.StatusSeeOther)

}

// Site

func dashboard(w http.ResponseWriter, req *http.Request) {

	// Zczytywanie ciastka o nazwie "id"
	cookie, err := req.Cookie("id")
	if err != nil {
		fmt.Println("er")
	}

	// Sprawdzanie wartości cookie - ciastka
	if cookie.Value != "" {
		fmt.Printf(cookie.Value)
	} else {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}

	c := &http.Cookie {
		Name: "id",
		MaxAge: -1,
		HttpOnly: true,
	}

	http.SetCookie(w, c)

	req.AddCookie(c)

	dat, err := ioutil.ReadFile("dashboard-analytics.html")

	if err != nil {
		fmt.Println(err)
	}

	x := strings.Index(string(dat), "<!-- specjalny komentarz -->");
	y := strings.Index(string(dat), "^");

	fmt.Println(y)

	var s string;
	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Jest blad i chuj");
	}

	result := db.Table("users").Where("ID = ?", cookie.Value).Find(&User{})
	var item User

	result.First(&item) // pierwszy item który spełnia wymagania ( name = ? )
	name := item.Name

	var records []Websites
	db.Table(name + "_websites").Find(&records)

	for n, i := range string(dat) {
		s += string(i);
		if n == x+28 {
			for _, x := range records {
				s += "<li class=\"list-item\">"+ x.Url +"</li>"
			}
		}
		// Kiedy n bedzie na poziomie value w form'ie
		if n == y-1 {
			s += string(name)
		}
	}

	//s = strings.ReplaceAll(s, "0x1c", "value.push("");value.push(\"\");value.push(\"\");value.push(\"\");value.push(\"\");value.push(\"\");value.push(\"\");")

	fmt.Fprintf(w, string(s))
}

func addwebsite(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil { // Parsowanie inputów
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	websitename := req.FormValue("addwebsite")
	username := req.FormValue("name")
	var masno string;

	// Usuwanie ostatniej litery z nazwy (masno) gdzie masno to jest nazwa usera ukrócona o 1 literę
	for nr, obj := range username {
		if nr < (len(username) - 2) {
			masno += string(obj);
		}
	}

	const (
		end = "_38_"
		dot = "_46_"
		slash = "_47_"
		doubledot = "_58_"
		questionMark = "_63_"
		equal = "_61_"
		minus = "_45_"
		plus = "_43_"
	)

	var url string = websitename

	url = strings.ReplaceAll(url, ".", dot)
	url = strings.ReplaceAll(url, "&", end)
	url = strings.ReplaceAll(url, "/", slash)
	url = strings.ReplaceAll(url, ":", doubledot)
	url = strings.ReplaceAll(url, "?", questionMark)
	url = strings.ReplaceAll(url, "=", equal)
	url = strings.ReplaceAll(url, "-", minus)
	url = strings.ReplaceAll(url, "+", plus)

	fmt.Println(url)

	db, err := gorm.Open(sqlite.Open("db.db" ), &gorm.Config{})
	if err != nil {
		fmt.Println("Wyjebalo bleda spierdalaj")
	}

	db.Migrator().CreateTable(masno + "_website_" + url)
	db.Migrator().AddColumn(&Website{}, "oj tak tak")
	pola := []string{" "}

	for nr := 0; nr < 8; nr++ {
		db.Migrator().AddColumn(&Website{}, "oj tak tak")
	}

	db.Table(masno + "_website_" + url).Select("ID", "Service", "Rang", "Days", "SSL", "Up", "Response", "Down").Create(&Website{ID: 1, Service: true, Rang: 2, Days: 7, SSL: true, Up: 2, Response: 7.35312321, Down: 3})
	fmt.Println("Wysłano")
	http.Redirect(w, req, "/dashboard", http.StatusSeeOther)

	//cos := beeep.Notify("Page status", "Page status is up", "");
}

// Główny serwer

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/register", register)
	http.HandleFunc("/dashboard", dashboard) // dashboard-analytics.html
	http.HandleFunc("/addwebsite", addwebsite)
	http.Handle("/", http.FileServer(http.Dir("C:/Users/Ninja/go/src/test2/")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Nie udało się uruchomić")
	}
}