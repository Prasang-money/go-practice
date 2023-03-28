package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Prasang-money/snippetbox/pkg/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check if the current request URL path exactly matches with "/". If it does not matches, we are sending 404 page not found error to the client.
	// importantly return after http.NotFound() otherwise function will keep wexecuting till the end.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.ServerError(w, err)
		return
	}
	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string{
	// 	"ui/html/home.page.tmpl",
	// 	"ui/html/base.layout.tmpl",
	// 	"ui/html/footer.partial.tmpl",
	// }
	// temp, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.ServerError(w, err)
	// 	return
	// }

	// err = temp.Execute(w, nil)
	// if err != nil {
	// 	app.ServerError(w, err)
	// 	return
	// }

}
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d.......", id)
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.ServerError(w, err)
		return
	}
	// Write the snippet data as a plain-text HTTP response body.
	fmt.Fprintf(w, "%v", s)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"
	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
		w.Write([]byte("Create a new snippet"))
	}

	// Redirect the user to the relevant page for the snippet.
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
