package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("/", home)
	serverMux.HandleFunc("/snippet", showSnippet)
	serverMux.HandleFunc("/snippet/create", createSnippet)

	con, err := net.Listen("tcp", ":8085")
	if err != nil {
		panic(err)
	}
	log.Println("starting server on port 8085")
	if err := http.Serve(con, serverMux); err != nil {
		panic(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	// check if the current request URL path exactly matches with "/". If it does not matches, we are sending 404 page not found error to the client.
	// importantly return after http.NotFound() otherwise function will keep wexecuting till the end.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from snippetbox"))
}
func showSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	fmt.Fprintf(w, "Display a specific snippet with ID %d.......", id)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Display as specific snippet...."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet"))
}
