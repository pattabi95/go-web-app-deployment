package main

import (
    "html/template"
    "log"
    "net/http"
)

type Experience struct {
    Role       string
    Company    string
    Location   string
    Start      string
    End        string
    Highlights []string
}

type SkillCategory struct {
    Category string
    Skills   []string
}

type PortfolioData struct {
    Name       string
    Title      string
    Summary    string
    Email      string
    LinkedIn   string
    Phone      string
    Page       string
    Experience []Experience
    Skills     []SkillCategory
}

func main() {
    tmpl := template.Must(template.ParseFiles("templates/layout.html"))

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := getPortfolioData()
        data.Page = "home"
        if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
            log.Println("Error rendering home:", err)
        }
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        data := getPortfolioData()
        data.Page = "about"
        if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
            log.Println("Error rendering about:", err)
        }
    })

    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        data := getPortfolioData()
        data.Page = "contact"
        if err := tmpl.ExecuteTemplate(w, "layout", data); err != nil {
            log.Println("Error rendering contact:", err)
        }
    })

    log.Println("Server running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getPortfolioData() PortfolioData {
    return PortfolioData{
        Name:     "Pattabiraman V K",
        Title:    "Cloud & DevOps Engineer",
        Summary:  "Cloud & DevOps Engineer | AWS Certified | CI/CD | Docker | Kubernetes | Terraform | Ex-Account Manager with Vendor & Project Leadership",
        Email:    "pattabimanu@gmail.com",
        LinkedIn: "linkedin.com/in/pattabi-raman",
        Phone:    "+91 8939608210",
        Experience: []Experience{
            {
                Role:    "Cloud/DevOps Engineer",
                Company: "Amazon",
                Location: "Chennai, IN",
                Start:   "Jan 2023",
                End:     "Present",
                Highlights: []string{
                    "Designed and implemented CI/CD pipelines using Jenkins and GitHub Actions",
                    "Containerized applications with Docker and deployed to EKS using Terraform",
                    "Implemented Kubernetes ingress, autoscaling and observability using Prometheus & Grafana",
                    "Developed reusable Terraform modules and maintained infrastructure as code",
                    "Mentored teams in Git workflows and DevOps best practices",
                },
            },
            {
                Role:    "Catalog Lead",
                Company: "Amazon",
                Location: "Chennai, IN",
                Start:   "Jan 2021",
                End:     "Dec 2022",
                Highlights: []string{
                    "Led onboarding of Apple NPI for Canada marketplace with notable impact",
                    "Streamlined vendor self‑service tooling, saving 13+ FTE",
                    "Managed audits across 18k+ ASINs for event readiness",
                },
            },
        },
        Skills: []SkillCategory{
            {"Cloud Technologies", []string{"AWS (EC2, VPC, IAM, EKS, ALB)"}},
            {"CI/CD", []string{"GitHub Actions", "Jenkins", "ArgoCD"}},
            {"Containers & Orchestration", []string{"Docker", "Kubernetes"}},
            {"IaC & Config Management", []string{"Terraform", "Ansible"}},
            {"Monitoring & Observability", []string{"Prometheus", "Grafana"}},
            {"Scripting & Tools", []string{"Shell", "Python", "Git"}},
        },
    }
}
