package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{} {
	"Title": "Personal Web",
	"IsLogin": true,
}

type dataReceive struct {
	// ID int
	Projectname string
	Description string
	// Technologies []string
	// Startdate string
	// Enddate string
	// Duration string
}

// Nanti si variable dataSubmit ini bakal di isi sama value dari function di addProject
var dataSubmit = []dataReceive{

}

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", helloWorld).Methods("GET")
	route.HandleFunc("/home", home).Methods("GET").Name("home")
	route.HandleFunc("/project", formProject).Methods("GET")
	route.HandleFunc("/project/{id}", detailProject).Methods("GET")
	route.HandleFunc("/project", addProject).Methods("POST")
	route.HandleFunc("/deleteProject/{id}", deleteProject).Methods("GET")
	route.HandleFunc("/contact", contactMe).Methods("GET")

	fmt.Println("Server is running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello world!"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	respData := map[string]interface{} {
		"Projects": dataSubmit,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/blog-page.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	projectname := r.PostForm.Get("name")
	// startDate := r.PostForm.Get("start-date")
	// endDate := r.PostForm.Get("end-date")
	description := r.PostForm.Get("description")

	// fmt.Println("Project Name :" + r.PostForm.Get("name"))
	// fmt.Println("Description :" + r.PostForm.Get("description"))

	var newData = dataReceive{
		Projectname: projectname,
		Description: description,
		// Technologies: technologies,
		// Startdate: startDate,
		// Enddate: endDate,
		// Duration: duration,
	} 

	dataSubmit = append(dataSubmit, newData)

	fmt.Println(dataSubmit)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	fmt.Println(id)

	dataSubmit = append(dataSubmit[:id], dataSubmit[id+1:]...)

	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var tmpl, err = template.ParseFiles("views/blog.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	resp := map[string]interface{} {
		"Data": Data,
		"Id": id,
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, resp)
}

func contactMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("views/contact.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, Data)
}