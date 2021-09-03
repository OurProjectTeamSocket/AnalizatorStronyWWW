package adds
//
//import (
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//	"log"
//)
//
//// Get
//
//func getUser(id string) (*User, error) {
//	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	result := db.Table("users").Last(&User{})
//
//	var item User
//
//	result.First(&item)
//
//	return &item, nil
//}
//
//func getWebsites(id string, username string) (*Websites, error) {
//	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	result := db.Table(username + "_websites").Last(&User{})
//
//	var item Websites
//
//	result.First(&item)
//
//	return &item, nil
//}
//
//func getWebsite( username string, url string) (*Website, error) {
//	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//
//	result := db.Table(username+"_website_"+Encode(url)).Last(&Website{})
//
//	var item Website
//
//	result.First(&item)
//
//	return &item, nil
//}
//
//// Put
//
//func putUser(name string, email string, password string) error {
//	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
//	if err != nil {
//		return err
//	}
//
//	result := db.Table("users").Last(&User{})
//
//	var item User
//
//	result.First(&item)
//
//	hash, err := Hash(password)
//	if err != nil {
//		return err
//	}
//
//	result = db.Table("users").Select("ID", "Name", "Email", "Password").Create(User{ID: item.ID+1, Name: name, Email: email, Password: hash})
//
//	log.Printf("ID: %v, Name: %v, Email: %v, Password: %v\n", item.ID+1, name, email, password)
//
//	return nil
//
//}
//
//func putWebsites(username string, url string) error {
//	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
//	if err != nil {
//		return err
//	}
//
//	var item Websites
//
//	db.Table(username + "_websites").Last(&item)
//
//	db.Table(username + "_websites").Select("ID", "url").Create(&Websites{ID: item.ID+1, Url: Encode(url)})
//
//	log.Printf("ID: %v, Url: %v\n", item.ID+1, url)
//
//	return nil
//}
//
//func putWebsite( username string, url string, service bool, rang int, days int, ssl bool, up int, response float64, down int) error {
//	db, err := gorm.Open(sqlite.Open("db.db"), &gorm.Config{})
//	if err != nil {
//		return err
//	}
//
//	result := db.Table(username+"_website_"+Encode(url)).Select("ID").First(&Website{})
//
//	var item Website
//
//	result.First(&item)
//
//	db.Table(username+"_website_"+Encode(url)).Select("ID", "Service", "Rang", "Days", "SSL", "Up", "Response", "down").Create(&Website{ID: item.ID+1, Service: service, Rang: rang, Days: days, SSL: ssl, Up: up, Response: response, Down: down})
//
//	log.Printf("ID: %v, Service: %v, Rang: %v, Days: %v, SSL: %v, Up: %v, Response: %v, down: %v\n", item.ID+1, service, rang, days, ssl, up, response, down)
//
//	return nil
//}