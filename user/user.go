package user

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/user", userHandler)

	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	// [END setting_port]
}

// indexHandler responds to requests with our greeting.
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/user" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello World from user!")
}
