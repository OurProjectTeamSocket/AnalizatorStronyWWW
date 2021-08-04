package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { // Funckja wspomniana w linijce 20
	dat, err := ioutil.ReadFile("index.html") // wczytywanie HTML'a
	if err != nil { // Jezeli nie ma pliku
		fmt.Println("Con not load the file with website's content")
	}
	fmt.Fprintf(w, string(dat)) // Wysyłanie jako zawartości
}


func main() {
	http.HandleFunc("/", helloHandler) // Nasłuchiwanie na "ip/" i wykonywanie funkcji
	//http.Handle("/", http.FileServer(http.Dir("./"))) // Nasłuchiwanie na "ip/" i wyświetlanie pliku


	fmt.Printf("Starting server at port 8080\n") // Info że serwer działa
	if err := http.ListenAndServe(":8080", nil); err != nil { // Sprawdzanie czy sie serwer nie wywalił na twarz
		log.Fatal(err)
	}
}
