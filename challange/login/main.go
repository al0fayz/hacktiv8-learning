package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"text/template"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/login", LoginHandler)

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

type Biodata struct {
	Id        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
	Email     string
	Password  string
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

		//open and read file
		user, err := OpenAndReadFile()
		if err != nil {
			fmt.Println(err)
		}
		//search by nama
		for i := 0; i < len(user); i++ {
			if user[i].Email == email && user[i].Password == password {
				var credentialData = Biodata{
					Id:        user[i].Id,
					Nama:      user[i].Nama,
					Alamat:    user[i].Alamat,
					Pekerjaan: user[i].Pekerjaan,
					Alasan:    user[i].Alasan,
					Email:     user[i].Email,
					Password:  user[i].Password,
				}

				var filepath = path.Join("template", "user.html")
				tmpl, err := template.ParseFiles(filepath)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				if err := tmpl.Execute(w, credentialData); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		}
		var info = Info{
			IsError: true,
			Message: "Email / Password tidak benar",
		}
		var filepath = path.Join("template", "login.html")
		tmpl, err := template.ParseFiles(filepath)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, info); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

// function openfile dummy data
func OpenAndReadFile() ([]Biodata, error) {
	//open dummy file
	rootPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	pathFile := path.Join(rootPath, "/dummy.json")
	file, err := os.OpenFile(pathFile, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var students []Biodata
	err = json.Unmarshal(byteValue, &students)
	if err != nil {
		return nil, err
	}
	//return
	return students, nil
}
