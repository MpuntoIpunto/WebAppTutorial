package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

//Page is a stuct for web page
type Page struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
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
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
	log.Println(title + " was loaded!")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	log.Println(title + " was loaded!")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Body: []byte(body), Title: title}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	log.Println(title + " was loaded!")
}

func main() {
	/* First part of tutorial, just write to a file an load it again
	p1 := &Page{Title: "Test Page", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("Test Page")
	fmt.Println(string(p2.Body)) */

	//Second part of Tutorial, make a simple webserver
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
