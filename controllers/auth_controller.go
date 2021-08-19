package controllers

import (
	"fmt"
	"net/http"

	"github.com/kataras/i18n"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Databse template

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

// Security

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil { // Input parsing
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		// inputy
		name := req.FormValue("user-name")
		password := req.FormValue("user-password")

		db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
		if err != nil {
			log.WithFields(log.Fields{"function": "login"}).Info(i18n.Tr("en-US", "db.connection.error"))
		} else {
			log.WithFields(log.Fields{"function": "login"}).Info(i18n.Tr("en-US", "db.connection.success"))
		}

		result := db.Table("users").Where("name = ?", name).Find(&User{})

		var item User

		result.First(&item) // the first item that meets the requirements (name =?)

		if !CheckHash(password, item.Password) { // If the passwords are NOT consistent
			log.WithFields(log.Fields{"function": "login"}).Warning(i18n.Tr("en-US", "auth.login.bad_password", name))
			return
		}

		log.WithFields(log.Fields{"function": "login"}).Info(i18n.Tr("en-US", "auth.login.success", name))

		db = nil

		http.Redirect(w, req, "/", http.StatusSeeOther) // Transfer to Homepage
	})
}

func Register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		name := req.FormValue("user-name")
		email := req.FormValue("user-email")
		password := req.FormValue("user-password")

		db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
		if err != nil {
			log.WithFields(log.Fields{"function": "login"}).Error(i18n.Tr("en-US", "db.connection.error"))
		} else {
			log.WithFields(log.Fields{"function": "login"}).Info(i18n.Tr("en-US", "db.connection.success"))
		}

		result := db.Table("users").Where("email = ?", email).Find(&User{})

		if result.RowsAffected > 0 { // If search results appeared
			log.WithFields(log.Fields{"function": "register"}).Info(i18n.Tr("en-US", "auth.register.duplicate", email))
			return
		}

		var item User

		myuuid := uuid.NewV4()
		fmt.Println(myuuid) // just using to make deployment pass

		pass, _ := Hash(password) // Password hashing

		db.Table("users").Select("ID", "Name", "Email", "Password").Create(&User{ID: /* myuuid */ item.ID + 1 /* adding 1 to the highest ID in the datum */, Name: name, Email: email, Password: pass})

		db = nil

	})
}
