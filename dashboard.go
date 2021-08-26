package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"strings"
)

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

	dat, err := ioutil.ReadFile("Stronadruga/dashboard-analytics.html") // Wczytywanie pliku

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

	// Wyciąganie informacji z bazy danych

	var info Website
	db.Table("Syta2_website_opt_46_software").Find(&info)

	s = strings.ReplaceAll(s, "0xf1", fmt.Sprintf(`<div class="day">%v</div>`, info.Up)) // info.ID, info.Service, info.Rang, info.Days, info.SSL, info.Up, info.Response, info.Down
	s = strings.ReplaceAll(s, "0xc1", fmt.Sprintf("%v", info.Down))
	s = strings.ReplaceAll(s, "0xcf", fmt.Sprintf("{\n    country: \"USA\",\n    visits: %v\n  }", info.Down)) /* ,
	{
	country: "China",
		visits: %v*10
	},
	{
	country: "Japan",
		visits: %v
	},
	{
	country: "Germany",
		visits: %v
	},
	{
	country: "UK",
		visits: %v*10
	},
	{
	country: "France",
		visits: %v
	},
	{
	country: "India",
		visits: %v
	}, {
	country: "Mongolia",
		visits: %v
	}*/

	fmt.Fprintf(w, string(s)) // wysyłanie strony
}
