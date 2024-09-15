package servers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hitenpratap/go-web-server/middleware"
	"github.com/joho/godotenv"
)

func StartNetHTTPServer() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//Serve static content
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/", fs)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	// Apply middleware
	wrappedMux := middleware.LoggingMiddleware(mux)

	fmt.Println("Net HTTP server starting at http://localhost:" + port)
	http.ListenAndServe(":"+port, wrappedMux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello with Middleware!")
}
