package main

// https://pkg.go.dev/net/http
import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./src"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloFunc)
	http.HandleFunc("/form", formFunc)

	fmt.Printf("Server is running on port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func formFunc(w http.ResponseWriter, r *http.Request){
	// ParseForm parses the raw query from the URL and updates r.Form.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")

	// FormValue get the form data with attribute name and address
	name := r.FormValue("name");
	address:= r.FormValue("address");

	// Write back the response to the client
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n" , address)	
}


func helloFunc(w http.ResponseWriter, r *http.Request){
	fmt.Println(w)
	fmt.Println(r)

	if r.URL.Path != "/hello" {
		http.Error(w, "404  not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
}
