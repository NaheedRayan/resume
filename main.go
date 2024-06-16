package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func renderPDF(w http.ResponseWriter, r *http.Request) {
    
    file_path := "pdfs/naheed_rayhan_cv_1.pdf"
    
    // Open the PDF file
    file, err := os.Open(file_path)
    if err != nil {
        http.Error(w, "Failed to open PDF file", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Write the PDF file to the response writer
    w.Header().Set("Content-Type", "application/pdf")
    http.ServeFile(w, r, file_path)
}

func main() {
    // Define and parse the command-line flag for the server address
    listenAddr := flag.String("listenAddr", ":8080", "the server address")
    flag.Parse()

    // Create a new router
    router := mux.NewRouter()

    // Register the handler for the root URL
    router.HandleFunc("/", renderPDF).Methods("GET")

    // Start the server
    fmt.Printf("Server running at port http://localhost%s", *listenAddr)
    log.Fatal(http.ListenAndServe(*listenAddr, router))
}
