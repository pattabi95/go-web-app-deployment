package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
)

type Profile struct {
    Name     string `json:"name"`
    Title    string `json:"title"`
    Location string `json:"location"`
    Summary  string `json:"summary"`
    LinkedIn string `json:"linkedin"`
}

var templates *template.Template
var profile Profile

func main() {
    // Parse templates with layout support
    templates = template.Must(template.ParseGlob("templates/*.html"))

    // Load profile data from JSON file
    file, err := os.Open("data/profile.json")
    if err != nil {
        log.Fatalf("Error opening profile.json: %v", err)
    }
    defer file.Close()
    if err := json.NewDecoder(file).Decode(&profile); err != nil {
        log.Fatalf("Error decoding profile json: %v", err)
    }

    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/about", aboutHandler)
    r.HandleFunc("/contact", contactHandler)
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    fmt.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "home.html", profile)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.html", profile)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "contact.html", profile)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}