package main

import (
    "fmt"
    "log"
	"html/template"
    "net/http"
    "github.com/gobuffalo/packr"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "[Webapp2] Hi there, is your name %s?", r.URL.Path[1:])
}

func main() {
	box := packr.NewBox("./templates")
    log.Printf("Running on port 8080")

    http.HandleFunc("/", rootHandler)

    staticTmpl := template.Must(template.New("staticTempl").Parse(box.String("layout.html")))    
	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		staticTmpl.Execute(w, data)
    })
    
    formsTmpl := template.Must(template.New("formsTmpl").Parse(box.String("forms.html")))    

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			formsTmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		_ = details

		formsTmpl.Execute(w, struct{ Success bool }{true})
	})    

	http.ListenAndServe(":8080", nil)
}