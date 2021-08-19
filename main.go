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

func register(w http.ResponseWriter, req *http.Request)  {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	name := req.FormValue("user-name") // Zmienna user-name ze strony logowania
	email := req.FormValue("user-email")  // Zmienna user-email ze strony logowania
	password := req.FormValue("user-password")  // Zmienna user-password ze strony logowania

	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(w, "Can not connect to database, sorry form problems. Please contact with administrator of this website")
	}

	result := db.Table("users").Where("name = ?", name).Find(&User{}) // w TABELI users GDZIE name rózna się zaawarośc name

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
		MaxAge: -1, // Ciasteczko nie istniejące
		HttpOnly: true,
	}

	http.SetCookie(w, c)

	req.AddCookie(c)

	dat, err := ioutil.ReadFile("dashboard-analytics.html") // Wczytywanie pliku

	if err != nil {
		fmt.Println(err)
	}

	x := strings.Index(string(dat), "<!-- specjalny komentarz -->") // Tam gdzie specialny komentarz to tam dodaje przyciski
	y := strings.Index(string(dat), "^") // Tu jest zmienna nazwa usera

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

	var records []Websites // Wiersze z DB ze stronami
	db.Table(name + "_websites").Find(&records) // Szukanie w nazwa_websites

	for n, i := range string(dat) {
		s += string(i);
		if n == x+28 { // Przyciski
			s += `<form onclick="action='/getInfoAboutSite'" method="POST">`
			for _, x := range records {
				s += `<li class="list-item"> <input type="submit" class="button" value="` + x.Url + `" name="website"> </li>`
			}
			s += `</form>`
		}
		// Kiedy n bedzie na poziomie value w form'ie
		if n == y-1 {
			s += string(name)
		}
	}

	s = strings.ReplaceAll(s, "0x1c", "value.push(" + strconv.Itoa(1) + ");\n" + // Eksportowanie danych stron WWW do wykresów
		"value.push(" + strconv.Itoa(2) + ");\n" +
		"value.push(" + strconv.Itoa(3) + ");\n" +
		"value.push(" + strconv.Itoa(4) + ");\n" +
		"value.push(" + strconv.Itoa(5) + ");\n" +
		"value.push(" + strconv.Itoa(6) + ");\n" +
		"value.push(" + strconv.Itoa(7) + ");")

	fmt.Fprintf(w, string(s)) // wysyłanie strony
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
	for nr, obj := range username { // usuwanie jednej litery z Go
		if nr < (len(username) - 1) {
			masno += string(obj);
		}
	}

	const ( // Kodowanie znaków specialnych
		end = "_38_"
		dot = "_46_"
		slash = "_47_"
		doubledot = "_58_"
		questionMark = "_63_"
		equal = "_61_"
		minus = "_45_"
		plus = "_43_"
	)

	var url string = websitename // URL

	//Zmienianie znaków zpecialnych na nasze kodowanie
	url = strings.ReplaceAll(url, ".", dot)
	url = strings.ReplaceAll(url, "&", end)
	url = strings.ReplaceAll(url, "/", slash)
	url = strings.ReplaceAll(url, ":", doubledot)
	url = strings.ReplaceAll(url, "?", questionMark)
	url = strings.ReplaceAll(url, "=", equal)
	url = strings.ReplaceAll(url, "-", minus)
	url = strings.ReplaceAll(url, "+", plus)

	fmt.Println(url) // wyświetlanie tego

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
	fmt.Println("Wysłano")

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