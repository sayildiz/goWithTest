package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// io.Writer interface has one method Write()
// os.Stdout and bytes.Buffer both have Write() method -> they implement the interface
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "word")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
