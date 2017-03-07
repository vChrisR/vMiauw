package main

import (
	"html/template"
	"net/http"
	"os"
	"vMiauw/config"

	"github.com/gorilla/mux"
)

type tContent struct {
	Title         string
	Head          string
	Message       string
	InstanceID    string
	InstanceIndex string
}

var c tContent

func vMiauw(w http.ResponseWriter, r *http.Request) {
	//t, err := template.New("vmtemplate").Parse("{{.Message}}")
	t, err := template.New("vmtemplate").ParseFiles("index.tmpl")
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "index.tmpl", c)
}

func main() {
	c.Message = "CatFoundry Demo"
	c.InstanceID = os.Getenv("CF_INSTANCE_GUID")
	c.Title = "vMiauw"
	c.Head = "Hello VMUG"
	c.InstanceIndex = os.Getenv("CF_INSTANCE_INDEX")

	router := mux.NewRouter()

	router.HandleFunc("/", vMiauw)

	http.Handle("/", router)

	http.ListenAndServe(":"+config.GetPort(), nil)
}
