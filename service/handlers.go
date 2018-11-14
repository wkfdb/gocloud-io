package service

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/unrolled/render"
)

func apiTestHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		cao := rand.Intn(10000)
		formatter.JSON(w, http.StatusOK, struct {
			Randnum int
		}{Randnum: cao})
	}
}

func homeHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		template := template.Must(template.New("index.html").ParseFiles("./templates/index.html"))
		_ = template.Execute(w, struct {
			ID      string
			Content string
		}{ID: "8675309", Content: "Hello from Go!"})
	}
}

func checkform(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := template.HTMLEscapeString(r.Form.Get("username"))
	password := template.HTMLEscapeString(r.Form.Get("password"))
	cao := strconv.Itoa(rand.Intn(100000))
	t := template.Must(template.New("signin.html").ParseFiles("./templates/signin.html"))
	err := t.Execute(w, struct {
		Username string
		Password string
		Randnum  string
	}{Username: username, Password: password, Randnum: cao})
	if err != nil {
		panic(err)
	}
}

func inDevelopmentHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("This is in development."))
}
