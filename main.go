package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"
)

const PORT = "8080"

var quotes = []string{
	"The web is more a social creation than a technical one. I designed it for a social effect—to help people work together. — Tim Berners-Lee",
	"Content precedes design. Design in the absence of content is not design, it’s decoration. — Jeffrey Zeldman",
	"The best websites are ones that give the users what they want, while at the same time meeting the needs of the business. — Steve Krug",
	"A website without visitors is like a ship lost in the horizon. — Dr. Christopher Dayagdag",
	"Websites should look good from the inside and out. — Paul Cookson",
	"Web design is not just about creating pretty layouts. It’s about understanding the marketing challenge behind your business. — Mohamed Saad",
	"The web is like a dominatrix. Everywhere I turn, I see little buttons ordering me to submit. — Nytwind",
	"Your website is the center of your digital ecosystem. Like a brick-and-mortar location, the experience matters once a customer enters. — Leland Dieno",
	"The great thing about web design is that you can do it on your own. All you need is a computer, an internet connection, and a lot of determination. — Mike Sullivan",
	"The key to making great websites is to not make them feel like websites. — Nick Disabato",
}

//go:embed index.html
var index embed.FS

func main() {
	indexFS, _ := fs.Sub(index, ".")
	http.Handle("/", http.FileServer(http.FS(indexFS)))
	http.HandleFunc("/quotes", eventsHandler)

	log.Println("Server listening on http://localhost:" + PORT)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal("error starting server on port " + PORT)
	}
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")

	for _, quote := range quotes {
		fmt.Fprintf(w, "data: %s\n\n", quote)
		w.(http.Flusher).Flush()
		time.Sleep(5 * time.Second)
	}
}
