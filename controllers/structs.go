package adds

// Szablon do bazy danych

type User struct {
	ID string
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