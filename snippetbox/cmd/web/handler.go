package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if the current request URL path exactly matches with "/". If it does not matches, we are sending 404 page not found error to the client.
	// importantly return after http.NotFound() otherwise function will keep wexecuting till the end.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"ui/html/home.page.tmpl",
		"ui/html/base.layout.tmpl",
		"ui/html/footer.partial.tmpl",
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "internal server error", 500)
		return
	}

	err = temp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}

}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d.......", id)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Display as specific snippet...."))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
