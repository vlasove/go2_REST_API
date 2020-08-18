package main

import (
	"fmt"
	"log"
	"net/http"
)

//HomePage function .....
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to Home page of ULTRA RESTAPI!!!!</h1>")
	fmt.Println("Hint: HomePage now linking....")
}

//GracefulShutdown function ....
func GracefulShutdown() {
	fmt.Println("REST API V1.0 shutting down....")
}

func main() {
	defer GracefulShutdown()
	fmt.Println("REST API V1.0 worked....")

	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

https://github.com/vlasove/go2_REST_API
