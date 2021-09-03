package adds

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func Register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := req.FormValue("user-name")         // Zmienna user-name ze strony logowania
		email := req.FormValue("user-email")       // Zmienna user-email ze strony logowania
		password := req.FormValue("user-password") // Zmienna user-password ze strony logowania

		if name == "" || password == "" || email == "" {
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}

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

		uid := uuid.New()
		pass, _ := Hash(password) // Haszowanie hasła
		db.Table("users").Select("ID", "Name", "Email", "Password").Create(&User{ID: uid.String(), Name: name, Email: email, Password: pass})

		db = nil

		http.Redirect(w, req, "/", http.StatusSeeOther)
	})
}
