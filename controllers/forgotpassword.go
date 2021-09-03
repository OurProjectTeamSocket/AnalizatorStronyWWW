package adds

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func ForgotPassword() http.HandlerFunc  {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		email := req.FormValue("user-email")
		if email == "" {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}

		db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
		if err != nil {
			log.Println(err)
		}

		var item User
		db.Table("users").Select("Email").First(&item)
		if item.Email == "" {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}

		log.Println("Hasło zostało pomyślnie zmienione!")
	})
}
