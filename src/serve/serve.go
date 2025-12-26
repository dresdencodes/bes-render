package serve

import (
	"os"
	"log"
	"net/http"
)

// next content to serve
var serveContent string
// next html
func NextHTML(html string) {
	_ = os.WriteFile("tmp/output.html", []byte(html), 0644)
	serveContent = html
}

// get url
func Run() {

	// validate server already running
	server := &http.Server{Addr: ":11111"}

	// handle func 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(serveContent))
	})

	log.Println("Serving on http://localhost:11111")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}