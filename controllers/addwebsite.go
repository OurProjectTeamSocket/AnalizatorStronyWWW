package adds

import (
	"encoding/hex"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func Addwebsite() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		fmt.Println("Work")

		cookie, err := req.Cookie("code")
		if err != nil {
			fmt.Println("er")
		}

		// Sprawdzanie wartości cookie - ciastka
		if cookie.Value != "" {
			fmt.Printf("Cookie: ", cookie.Value, "\n")
		} else {
			//http.Redirect(w, req, "/", http.StatusSeeOther)
		}

		if err := req.ParseForm(); err != nil { // Parsowanie inputów
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}





		websitename := req.FormValue("name")

		url := Encode(websitename)





		db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
		if err != nil {
			fmt.Println("SQL error connection")
		}





		str, err := hex.DecodeString(cookie.Value)
		if err != nil {
			fmt.Println(err)
		}





		var id_website Website

		db.Table(string(str) + "_websites").Last(&id_website) // Branie ID

		db.Table(string(str) + "_websites").Create(&Websites{ID: id_website.ID + 1, Url: websitename}) // Dodawanie strony do listy stron urzytkowanika

		fmt.Println("URL: ", url)

		db.Exec("CREATE TABLE " + string(str) + "_website_" + url + "(" + // Tworzenie DB da strony danego urzytkownika
			"ID INTEGER," +
			"Service INTEGER," +
			"Rang INTEGER," +
			"Days INTEGER," +
			"SSL INTEGER," +
			"Up INTEGER," +
			"Response DOUBLE," +
			"Down INTEGER" +
			")")

		db.Table(string(str)+"_website_"+url).Select("ID", "Service", "Rang", "Days", "SSL", "Up", "Response", "Down").Create(&Website{ID: 1, Service: true, Rang: 2, Days: 7, SSL: true, Up: 2, Response: 7.35312321, Down: 3}) // Podstawowe wartości przy dodawaniu strony

		//http.Redirect(w, req, "/dashboard", http.StatusSeeOther)

		//cos := beeep.Notify("Page status", "Page status is up", "");
	})
}
