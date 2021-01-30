package main

import "text/template"

// TODO: Serve this up on nginx

// how do I refresh?
func main() {
	tmpl := template.Must(template.ParseFiles("resources/html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// add chat messages
		// tmpl.Execute(w, ...)
	})
	http.ListenAndSerive(":80", nil)
}
