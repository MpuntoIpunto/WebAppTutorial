package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Page is a stuct for web page
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	log.Println(title + " was loaded!")
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
	/* First part of tutorial, just write to a file an load it again
	p1 := &Page{Title: "Test Page", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("Test Page")
	fmt.Println(string(p2.Body)) */

	//Second part of Tutorial, make a simple webserver
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
