package main
 import (
	"fmt"
	"net/http"
	"log"
 )

 func main(){
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
 }

 func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go-component!")
 }

