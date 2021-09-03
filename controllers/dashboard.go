package adds

import (
	"encoding/hex"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

func Dashboard() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {





		// Cookie

		// Zczytywanie ciastka o nazwie "id"
		cookie, err := req.Cookie("code")
		if err != nil {
			fmt.Println("er")
		}

		// Sprawdzanie wartości cookie - ciastka
		if cookie.Value != "" {
			fmt.Printf("Cookie: ", cookie.Value, "\n")
		} else {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}





		// Loading website

		dat, err := ioutil.ReadFile("./Website/dashboard.html") // Wczytywanie pliku

		if err != nil {
			fmt.Println(err)
		}





		// SQL Drivers

		db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
		}





		// HEX'a -> ASCII

		str, err := hex.DecodeString(cookie.Value)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Decoded name: ", string(str))





		// Downloading records form database

		var records []Websites
		db.Table(string(str) + "_websites").Find(&records)





		// Modyfing HTML code

		var urls string

		for _, x := range records {
			urls += `<li class="nav-item"><a href="#"><i class="fa fa-rocket"></i>` + x.Url + `</a></li>`
		}

		output := strings.ReplaceAll(string(dat), "kutasmaładnepałączki", urls)

		fmt.Fprintf(w, output)

	})

}
