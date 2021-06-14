package main

import (
	"fmt"
	"net/http"
)

const (
	aboutHTML = `<html>
	<body>
		<h1>About</h1>
		<h3>This text contains pen</h3>
		<h3>This text also contains mug</h3>
	</body>
	</html>`
	contactHTML = `<html>
	<body>
		<h1>Contact</h1>
		<h3>This text contains pen</h3>
		<h3>This text also contains mug</h3>
	</body>
	</html>`
)

func main() {
	http.HandleFunc("/about.html", func(res http.ResponseWriter, req *http.Request) {
		printReqHeaders(req)
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, aboutHTML)
	})
	http.HandleFunc("/admin/contact.html", func(res http.ResponseWriter, req *http.Request) {
		printReqHeaders(req)
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, contactHTML)
	})
	http.ListenAndServe(":5000", nil)
}

func printReqHeaders(r *http.Request) {
	for name, values := range r.Header {
		// Loop over all values for the name.
		for _, value := range values {
			fmt.Println(name, value)
		}
	}
}
