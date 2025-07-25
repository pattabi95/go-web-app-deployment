package main

import (
    "html/template"
    "net/http"
)

type Profile struct {
    Name     string
    Title    string
    Location string
    Summary  string
    LinkedIn string
}

func main() {
    tmpl := template.Must(template.ParseFiles("templates/profile.html"))

    profile := Profile{
        Name:     "Pattabi Raman",
        Title:    "Cloud/DevOps Engineer",
        Location: "India-Tamil Nadu-Chennai",
        Summary:  "Cloud & DevOps Engineer | AWS Certified | CI/CD | Docker | Kubernetes | Terraform | Ex-Account Manager with Vendor & Project Leadership",
        LinkedIn: "https://www.linkedin.com/in/pattabi-raman-805b47b8/",
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, profile)
    })

    http.ListenAndServe(":8080", nil)
}
