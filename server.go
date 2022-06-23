package main

import (
	"log"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(`
		<html>
			<head>
				<title>chat</title>
			</head>
		</html>
		<body>
			let's chat!
		</body>
	`))
	// fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	// start server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
