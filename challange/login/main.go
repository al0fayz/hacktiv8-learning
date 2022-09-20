package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func main() {
	http.HandleFunc("/", HomeHandler)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("images"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

type Info struct {
	IsError bool
	Message string
}

// halaman login
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var filepath = path.Join("template", "login.html")
		var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var info = Info{
			IsError: false,
			Message: "",
		}

		err = tmpl.Execute(w, info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

// handle login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var email = r.FormValue("email")
		var password = r.Form.Get("password")

		var data = map[string]string{"email": email, "password": password}

		var tmpl = template.Must(template.New("result").ParseFiles("view.html"))

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}
