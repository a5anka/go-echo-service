// (This is the same Go code we already have for the echo microservice)
package main

import (
    "fmt"
    "log"
    "net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
    prefix := "Echoed: " 
    inputString := r.URL.Query().Get("input")

    if inputString == "" {
        http.Error(w, "Missing 'input' query parameter", http.StatusBadRequest)
        return
    }

    echoedString := prefix + inputString
    fmt.Fprintln(w, echoedString)
}

func main() {
    http.HandleFunc("/echo", echoHandler)

    fmt.Println("Echo microservice listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
