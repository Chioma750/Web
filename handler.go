package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Checking for the route path
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page not found")
		return
	}

	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("text.html")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid requeest")
			return
		}
		tmpl.Execute(w, nil)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid method")
		return
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		tmpl, err := template.ParseFiles("result.html")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid request")
			return
		}
		Reader := r.FormValue("fname")
		Read := r.FormValue("lname")

	
		Result := map[string]interface{} {
			"Reader": Reader,
			"Read": Read,
		}
		tmpl.Execute(w, Result)
		fmt.Fprintf(w, Reader + " ")
		fmt.Fprintf(w, Read)


	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid method")
		return
	}
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/SERVER", handle)
	fmt.Println("server is running at http://localhost:8888")
	http.ListenAndServe(":8888", nil)
}
